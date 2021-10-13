package main

import (
	"fmt"
	"sync/atomic"
)

type food struct {
	name string
}

func main() {
	var v atomic.Value

	set(v, nil)

	f := v.Load()
	fmt.Println(f)

	fmt.Println(f.(*food))
}

func set(v atomic.Value, f *food) {
	v.Store(f)
}
