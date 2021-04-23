package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("----Creating a new Person----")
			return new(Person)
		},
	}
}

func main() {
	initPool()

	p := pool.Get().(*Person)
	fmt.Println("First get person:", p)

	p.Name = "jack"
	fmt.Printf("Set person name")

	pool.Put(p)

	fmt.Println("Get persion", pool.Get().(*Person))
	fmt.Println("Get persion", pool.Get().(*Person))
	fmt.Println("Get persion", pool.Get().(*Person))
}
