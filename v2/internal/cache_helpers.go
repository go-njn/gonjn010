package internal

import "time"

const (
	keyIsEmptyErrText      = "key is empty"
	keyNotFoundErrTemplate = "key not found, key = %q"
	valueIsEmptyErrText    = "value is empty"
	itemExpiredErrTemplate = "item expired, key = %q"
	ttlTooShortErrTemplate = "life %q time is too short, min = %q"
	ttlTooHighErrTemplate  = "life time %q is too high, max = %q"
)

type cacheItem struct {
	value       any //todo: consider generic type
	whenExpired time.Time
}

func getNow() time.Time {
	return time.Now() //todo: inject time source
}

func getExpiration(ttl time.Duration) time.Time {
	return getNow().Add(ttl)
}

func isExpired(item cacheItem) bool {
	return getNow().After(item.whenExpired)
}
