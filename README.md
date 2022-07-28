# Cache

___

In-memory cache implementation, Go package

# V2 with time-to-live for value item
```shell
go get -u -v  github.com/go-njn/gonjn010/v2
```
```go
...
  Set(key string, value any, ttl time.Duration)
  Get(key string) (any, error)
  Delete(key string) error
...
```
```go
package main

import (
	cacheV1 "github.com/go-njn/gonjn010/cache"
	cacheV2 "github.com/go-njn/gonjn010/v2/pkg/cache"
	"time"
)

func main() {
	println("v1 version:", cacheV1.Version())
	v1 := cacheV1.New()
	v1.Set("user1", 42)

	println("v2 version:", cacheV2.Version())
	v2 := cacheV2.New()
	v2.Set("user1 v2", 42, 5 * time.Second)
	item, err := v2.Get("user1 v2")
	if err != nil {
		println("error: ", err.Error())
	} else {
		println("got item: ", item.(int))
	}

	println("\nitem expiration test")
	println("waiting 6 sec, item must expire...")
	time.Sleep(6 * time.Second)
	item, err = v2.Get("user1 v2")
	if err != nil {
		println("error: ", err.Error())
	} else {
		println("got item: ", item.(int))
	}

	println("\ndelete item test")
	v2.Set("user2 v2", 43, 5*time.Second)
	item, err = v2.Get("user2 v2")
	if err != nil {
		println("error: ", err.Error())
	} else {
		println("got item: ", item.(int))
	}

	println("removing...")
	v2.Delete("user2 v2")
	item, err = v2.Get("user2 v2")
	if err != nil {
		println("error: ", err.Error())
	} else {
		println("got item: ", item.(int))
	}

}
```
```text
v1 version: v1.1.1
v2 version: v2.0.0
got item:  42

item expiration test
waiting 6 sec, item must expire...
error:  item expired, key = "user1 v2" 

delete item test
got item:  43
removing...
error:  key not found, key = "user2 v2"
```
___
# V1 preliminary implementation

```go
  Set(key string, value any) error
  Get(key string) (any, error)
  Delete(key string) error
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
v1.1.1
```
___
