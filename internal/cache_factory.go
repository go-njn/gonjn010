package internal

import (
	"time"
)

func New(defaultItemLifeTime ...time.Duration) *cacheImpl {
	lifeTime := defaultLifeTime
	if len(defaultItemLifeTime) > 0 {
		lifeTime = defaultItemLifeTime[0]
	}

	return &cacheImpl{
		lifeTime: lifeTime,
		storage:  make(map[string]cacheItem),
	}
}
