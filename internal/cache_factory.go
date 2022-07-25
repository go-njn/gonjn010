package internal

import (
	"time"
)

func New(defaultItemLifeTime ...time.Duration) *Cache {
	lifeTime := defaultLifeTime
	if len(defaultItemLifeTime) > 0 {
		lifeTime = defaultItemLifeTime[0]
	}

	return &Cache{
		lifeTime: lifeTime,
		storage:  make(map[string]cacheItem),
	}
}
