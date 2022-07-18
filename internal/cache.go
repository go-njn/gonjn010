package internal

import (
	"errors"
	"sync"
	"time"
)

type cacheImpl struct {
	sync     sync.RWMutex
	lifeTime time.Duration
	storage  map[string]cacheItem
}

func (c *cacheImpl) set(key string, value any, lifeTime time.Duration) error {
	c.sync.Lock()
	defer c.sync.Unlock()

	c.storage[key] = cacheItem{
		value:       value,
		whenExpired: getExpiration(lifeTime),
	}

	return nil
}

func (c *cacheImpl) Set(key string, value any) error {
	return c.set(key, value, c.lifeTime)
}

func (c *cacheImpl) Get(key string) (any, error) {
	c.sync.RLock()
	defer c.sync.RUnlock()

	item, found := c.storage[key]
	if !found {
		return nil, errors.New(keyNotFoundErrText)
	}

	if isExpired(item) {
		return nil, errors.New(itemExpiredErrText)
	}

	return item.value, nil
}

func (c *cacheImpl) Delete(key string) error {
	c.sync.Lock()
	defer c.sync.Unlock()

	delete(c.storage, key)

	return nil
}
