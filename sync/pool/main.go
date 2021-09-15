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
	pool.Put(p)
	fmt.Println("Set person jack")

	p2 := &Person{Name: "lucy"}
	pool.Put(p2)
	fmt.Println("Set person lucy")

	p3 := &Person{Name: "lilei"}
	pool.Put(p3)
	fmt.Println("Set person lilei")

	fmt.Println("Get persion", pool.Get().(*Person))
	fmt.Println("Get persion", pool.Get().(*Person))
	fmt.Println("Get persion", pool.Get().(*Person))
}
