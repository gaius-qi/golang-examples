package main

import (
	"bytes"
	"fmt"

	"github.com/golang/groupcache/lru"
)

func main() {
	cache := lru.New(2)
	buf1 := bytes.NewBuffer([]byte("hello"))
	buf2 := bytes.NewBuffer([]byte("work"))
	buf3 := bytes.NewBuffer([]byte("!"))

	cache.Add("p1", buf1)
	cache.Add("p2", buf2)

	v1, ok := cache.Get("p1")
	if ok {
		fmt.Printf("p1 %v\n", v1)
	}

	cache.Add("p3", buf3)

	fmt.Printf("cache length is %d\n", cache.Len())

	_, ok = cache.Get("p2")
	if !ok {
		fmt.Printf("p2 was evicted out\n")
	}
}
