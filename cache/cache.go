package cache

import (
	"github.com/go-njn/gonjn010/internal"
	"time"
)

type Cache internal.Cache

func Version() string {
	return internal.Version
}

func New(lifeTime ...time.Duration) Cache {
	return internal.New(lifeTime...)
}
