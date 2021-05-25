package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

func main() {
	addInt32()
	swapInt64()
	swapPointer()
	value()
}

func addInt32() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&n, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt32(&n)) // 1000
}

type T struct{ x int }

var pT *T

func swapInt64() {
	var n int64 = 123
	var old = atomic.SwapInt64(&n, 789)
	fmt.Println(n, old) // 789 123
	swapped := atomic.CompareAndSwapInt64(&n, 123, 456)
	fmt.Println(swapped) // false
	fmt.Println(n)       // 789
	swapped = atomic.CompareAndSwapInt64(&n, 789, 456)
	fmt.Println(swapped) // true
	fmt.Println(n)       // 456
}

func swapPointer() {
	var unsafePPT = (*unsafe.Pointer)(unsafe.Pointer(&pT))
	var ta, tb = T{1}, T{2}
	atomic.StorePointer(
		unsafePPT, unsafe.Pointer(&ta))
	fmt.Println(pT) // &{1}
	pa1 := (*T)(atomic.LoadPointer(unsafePPT))
	fmt.Println(pa1 == &ta) // true
	pa2 := atomic.SwapPointer(
		unsafePPT, unsafe.Pointer(&tb))
	fmt.Println((*T)(pa2) == &ta) // true
	fmt.Println(pT)               // &{2}
	b := atomic.CompareAndSwapPointer(
		unsafePPT, pa2, unsafe.Pointer(&tb))
	fmt.Println(b) // false
	b = atomic.CompareAndSwapPointer(
		unsafePPT, unsafe.Pointer(&tb), pa2)
	fmt.Println(b) // true
}

// go 1.17+
func value() {
	type T struct{ a, b, c int }
	var x = T{1, 2, 3}
	var y = T{4, 5, 6}
	var z = T{7, 8, 9}
	var v atomic.Value
	v.Store(x)
	fmt.Println(v) // {{1 2 3}}
	old := v.Swap(y)
	fmt.Println(v)       // {{4 5 6}}
	fmt.Println(old.(T)) // {1 2 3}
	swapped := v.CompareAndSwap(x, z)
	fmt.Println(swapped, v) // false {{4 5 6}}
	swapped = v.CompareAndSwap(y, z)
	fmt.Println(swapped, v) // true {{7 8 9}}
}
