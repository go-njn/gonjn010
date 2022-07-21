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
```shell
go get -u -v  github.com/go-njn/gonjn010
```
___

# example 1
[play](https://goplay.space/#2xH2YJBUk_B "goplay.space")
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
[play](https://goplay.space/#W9Wccxu89Hj "goplay.space")
```go
package main

import (
	"fmt"
	"time"

	memcache "github.com/go-njn/gonjn010/cache"
)

func main() {
	cache := memcache.New(15 * time.Minute)

	if err := cache.Set("userId", 42); err != nil {
		fmt.Println(err)
	}

	cache = memcache.New(3 * time.Minute)

	if err := cache.Set("userId", 42); err != nil {
		fmt.Println(err)
	}

	if value, err := cache.Get("userId"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	if err := cache.Delete("xxx-userId-xxx"); err != nil {
		fmt.Println(err)
	}

	if err := cache.Delete("userId"); err != nil {
		fmt.Println(err)
	}

	if userId, err := cache.Get("userId"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}

	fmt.Println(memcache.Version())
}
```
```text
life time "15:00" is too high, max = "05:00"
42
key not found, key = "xxx-userId-xxx"
key not found, key = "userId"
v1.1.0
```
___
