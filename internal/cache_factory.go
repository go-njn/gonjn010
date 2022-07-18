package internal

import (
	"github.com/go-njn/gonjn010/cache"
	"time"
)

func NewCache(defaultItemLifeTime ...time.Duration) cache.Cache {
	lifeTime := defaultLifeTime
	if len(defaultItemLifeTime) > 0 {
		lifeTime = defaultItemLifeTime[0]
	}

	return &cacheImpl{
		lifeTime: lifeTime,
		storage:  make(map[string]cacheItem),
	}
}
