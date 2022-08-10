package cache

import (
	"github.com/go-njn/gonjn010/v2/internal"
)

func Version() string {
	return internal.Version
}

func New() *internal.Cache {
	return internal.New()
}
