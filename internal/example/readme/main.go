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
	v2.Set("user1 v2", 42, 5*time.Second)
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
