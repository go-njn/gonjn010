package internal

import (
	"errors"
	"fmt"
	"time"
)

func validateKey(key string) error {
	if len(key) > 0 {
		return nil
	}

	return errors.New(keyIsEmptyErrText)
}

func validateValue(value any) error {
	if value != nil {
		return nil
	}

	return errors.New(valueIsEmptyErrText)
}

func validateItemLifeTime(ttl time.Duration) error {
	const minTTL = time.Second * 1 //todo: clarify business requirement for min\max TTL
	const maxTTL = time.Minute * 5
	const format = "04:05"
	if ttl < minTTL {
		return fmt.Errorf(ttlTooShortErrTemplate,
			formatDuration(ttl, format),
			formatDuration(minTTL, format))
	}

	if ttl > maxTTL {
		return fmt.Errorf(ttlTooHighErrTemplate,
			formatDuration(ttl, format),
			formatDuration(maxTTL, format))
	}

	return errors.New(valueIsEmptyErrText)
}

func getKeyNotFoundError(key string) error {
	return fmt.Errorf(keyNotFoundErrTemplate, key)
}

func getItemExpiredError(key string) error {
	return fmt.Errorf(itemExpiredErrTemplate, key)
}

func formatDuration(d time.Duration, format string) string {
	return time.Unix(0, 0).UTC().Add(d).Format(format)
}
