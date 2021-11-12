package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	m := make(map[string]struct{}, 1000000)
	s := make(map[string]struct{}, 1000000)
	h := make(map[string]struct{}, 1000000)
	fmt.Println(unsafe.Sizeof(m))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(h))
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Printf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}
