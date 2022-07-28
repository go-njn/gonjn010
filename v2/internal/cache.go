package internal

import (
	"sync"
	"time"
)

type Cache struct {
	locker  *sync.RWMutex
	storage map[string]cacheItem
}

func (c *Cache) Set(key string, value any, ttl time.Duration) error {
	if err := validateKey(key); err != nil {
		return err
	}

	if err := validateValue(value); err != nil {
		return err
	}

	if err := validateItemLifeTime(ttl); err != nil {
		return err
	}

	c.locker.Lock()
	defer c.locker.Unlock()

	c.storage[key] = cacheItem{
		value:       value,
		whenExpired: getExpiration(ttl),
	}

	return nil
}

func (c *Cache) Get(key string) (any, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	c.locker.RLock()
	item, found := c.storage[key]
	c.locker.RUnlock()

	if !found {
		return nil, getKeyNotFoundError(key)
	}

	if isExpired(item) {
		return nil, getItemExpiredError(key)
	}

	return item.value, nil
}

func (c *Cache) Delete(key string) error {
	if err := validateKey(key); err != nil {
		return err
	}

	c.locker.Lock()
	defer c.locker.Unlock()

	if _, found := c.storage[key]; !found {
		return getKeyNotFoundError(key)
	}

	delete(c.storage, key)

	return nil
}
