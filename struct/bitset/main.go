package main

import (
	"fmt"

	"github.com/willf/bitset"
)

func main() {
	b := bitset.New(2)
	fmt.Printf("empty bitset: %v\n", b.All())

	b.Set(0)
	fmt.Printf("0 bitset: %v\n", b.All())
	fmt.Printf("size bitset: %v\n", b.Count())

	b.Set(0).Set(1)
	fmt.Printf("0 1 bitset: %v\n", b.All())
}
