package internal

import "sync"

func New() *Cache {
	return &Cache{
		locker:  new(sync.RWMutex),
		storage: make(map[string]cacheItem),
	}
}
