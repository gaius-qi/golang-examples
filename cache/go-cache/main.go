package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type Animal struct {
	name   string
	number int
	height int
}

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Init cat
	cat := &Animal{
		name:   "jack",
		number: 1,
		height: 1,
	}

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("cat", cat, 3*time.Second)
	if x, found := c.Get("cat"); found {
		cat := x.(*Animal)
		fmt.Printf("name: %s, number: %d, height: %d\n", cat.name, cat.number, cat.height)
	}

	time.Sleep(2 * time.Second)
	x, duration, ok := c.GetWithExpiration("cat")
	fmt.Printf("expire: %t, duration: %s, value: %#v\n", ok, duration, x)

	x, found := c.Get("cat")
	fmt.Printf("found: %t, value: %#v, items: %#v\n", found, x, c.Items())
}
