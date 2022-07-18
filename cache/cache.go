package cache

import "github.com/go-njn/gonjn010/internal"

type Cache interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
}

func Version() string {
	return internal.Version
}

func NewCache() Cache {
	return nil
}