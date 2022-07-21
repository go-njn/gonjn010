package internal

import (
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
}

type cacheImpl struct {
	sync     sync.RWMutex
	lifeTime time.Duration
	storage  map[string]cacheItem
}

func (c *cacheImpl) set(key string, value any, lifeTime time.Duration) error {
	if err := validateKey(key); err != nil {
		return err
	}

	if err := validateValue(value); err != nil {
		return err
	}

	if err := validateItemLifeTime(lifeTime); err != nil {
		return err
	}

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
	if err := validateKey(key); err != nil {
		return nil, err
	}

	c.sync.RLock()
	defer c.sync.RUnlock()

	item, found := c.storage[key]
	if !found {
		return nil, getKeyNotFoundError(key)
	}

	if isExpired(item) {
		return nil, getItemExpiredError(key)
	}

	return item.value, nil
}

func (c *cacheImpl) Delete(key string) error {
	if err := validateKey(key); err != nil {
		return err
	}

	c.sync.Lock()
	defer c.sync.Unlock()

	if _, found := c.storage[key]; !found {
		return getKeyNotFoundError(key)
	}

	delete(c.storage, key)

	return nil
}
