package cache

import (
	"github.com/go-njn/gonjn010/internal"
	"time"
)

func Version() string {
	return internal.Version
}

func New(lifeTime ...time.Duration) *internal.Cache {
	return internal.New(lifeTime...)
}
