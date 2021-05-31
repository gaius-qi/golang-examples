package main

import "runtime"

func main() {
	go runtime.GC()
	go runtime.GC()
	runtime.GC()
	runtime.GC()
}
