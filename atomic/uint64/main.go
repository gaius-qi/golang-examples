package main

import (
	"fmt"

	"sync/atomic"
)

func main() {
	var n atomic.Uint64

	n.Store(42)
	fmt.Println(n.Load())
}
