package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

type Animal struct {
	Name   string
	Number int
	Height int
}

type Item struct {
	Object     interface{}
	Expiration int64
}

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Init cat
	cat := &Animal{
		Name:   "jack",
		Number: 1,
		Height: 1,
	}

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("cat", cat, 3*time.Second)
	if x, found := c.Get("cat"); found {
		cat := x.(*Animal)
		fmt.Printf("name: %s, number: %d, height: %d\n", cat.Name, cat.Number, cat.Height)
	}

	time.Sleep(2 * time.Second)
	x, duration, ok := c.GetWithExpiration("cat")
	fmt.Printf("expire: %t, duration: %s, value: %#v\n", ok, duration, x)

	x, found := c.Get("cat")
	fmt.Printf("found: %t, value: %#v, items: %#v\n", found, x, c.Items())

	if err := c.SaveFile("./cache"); err != nil {
		fmt.Println("save file err:", err)
		return
	}

	if err := c.LoadFile("./cache"); err != nil {
		fmt.Println("load file err:", err)
		return
	}
	fmt.Println(c.Items())

	fp, err := os.Open("./cache")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

	dec := gob.NewDecoder(fp)
	items := map[string]Item{}
	if err := dec.Decode(&items); err != nil {
		fmt.Println("decode err:", err)
		return
	}

	fmt.Printf("dec: %#v, items: %#v", dec, items)
}
