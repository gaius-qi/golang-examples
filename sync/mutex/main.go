package main

import (
	"fmt"
	"sync"
)

func main() {
	m := &sync.Mutex{}
	m.Lock()
	fmt.Println("ok")
	m.Lock()
	fmt.Println("ok")
	m.Lock()
	fmt.Println("ok")
	m.Lock()
	fmt.Println("ok")
	m.Lock()
	fmt.Println("ok")
	m.Lock()
	fmt.Println("ok")
}
