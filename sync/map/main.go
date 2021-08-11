package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}

	m.Store("foo", "bar1")
	v1, _ := m.Load("foo")
	fmt.Println("v1:", v1)
	v2, _ := m.LoadOrStore("foo", "bar2")
	fmt.Println("v2:", v2)
	m.Store("foo", "bar3")
	v3, _ := m.Load("foo")
	fmt.Println("v3:", v3)
}
