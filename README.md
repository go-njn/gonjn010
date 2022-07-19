# Cache

___

In-memory cache implementation, Go package

```go
type Cache interface {
Set(key string, value any) error
Get(key string) (any, error)
Delete(key string) error
}
```

___

# example 1

```go
package main

import (
	"fmt"
	"github.com/go-njn/gonjn010/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)
}
```
```text
42
<nil>
```

___

# example 2

```go
package main

import (
	"fmt"
	memcache "github.com/go-njn/gonjn010/cache"
	"time"
)

func main() {
	cache := memcache.New(15 * time.Minute)

	if err := cache.Set("userId", 42); err != nil {
		fmt.Println(err)
	}

	if userId, err := cache.Get("userId"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}

	cache.Delete("userId")
	if userId, err := cache.Get("userId"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}

	fmt.Println(memcache.Version())
}
```
```text
42
key not found
v0.0.9  
```
___