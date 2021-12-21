package main

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

func main() {
	b := bitset.BitSet{}

	b.Set(uint(10))
	b.Set(uint(100))
	b.Set(uint(1000))
	b.Set(uint(10))
	b.Set(uint(0))
	fmt.Println("count: ", b.Count())
	fmt.Println("len: ", b.Len())
	fmt.Println("binary storage size: ", b.BinaryStorageSize())

	b.ClearAll()
	fmt.Println("count: ", b.Count())
}
