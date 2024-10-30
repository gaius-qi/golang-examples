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
	b.Set(uint(1000))
	fmt.Println("count: ", b.Count())
	fmt.Println("len: ", b.Len())
	fmt.Println("binary storage size: ", b.BinaryStorageSize())

	b.Set(uint(0))
	b.Set(uint(10))
	fmt.Println("count: ", b.Count())
	fmt.Println("len: ", b.Len())
	fmt.Println("binary storage size: ", b.BinaryStorageSize())

	b.ClearAll()
	fmt.Println("count: ", b.Count())

	b1 := bitset.New(1000)
	b1.SetAll()
	fmt.Println("b1 length: ", b1.Len())
	fmt.Println("b1 count: ", b1.Count())

	b2 := bitset.New(1000)
	b2.Set(uint(10))
	fmt.Println("b2 length: ", b2.Len())
	fmt.Println("b2 count: ", b2.Count())
}
