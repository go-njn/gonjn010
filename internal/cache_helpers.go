package internal

import "time"

const (
	keyNotFoundErrText = "key not found"
	itemExpiredErrText = "item expired"
)

const (
	defaultLifeTime time.Duration = 5 * time.Minute
)

type cacheItem struct {
	value       any //todo: consider generic type
	whenExpired time.Time
}

func getNow() time.Time {
	return time.Now() //todo: inject time source
}

func getExpiration(duration time.Duration) time.Time {
	return getNow().Add(duration)
}

func isExpired(item cacheItem) bool {
	return getNow().After(item.whenExpired)
}
