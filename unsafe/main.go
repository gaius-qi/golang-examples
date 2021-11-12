package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"unsafe"
)

func main() {
	m := make([]map[string]struct{}, 1000000)
	s := make([]map[string]struct{}, 1000000)
	h := make([]map[string]struct{}, 1000000)
	fmt.Println(unsafe.Sizeof(m))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(h))
	fmt.Println(unsafe.Sizeof([]map[string]struct{}{}))

	d := []map[string]struct{}{}
	for i := 0; i < 1000000; i++ {
		d = append(d, map[string]struct{}{})
	}

	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Printf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
	fmt.Printf("%.2f MB", float64(rtm.HeapAlloc)/1024./1024.)
	fmt.Printf("%.2f MB", float64(rtm.Mallocs)/1024./1024.)
	fmt.Printf("%.2f MB", float64(rtm.MSpanInuse)/1024./1024.)

	_ = m

	fmt.Println("---------------")

	b := new(bytes.Buffer)
	gob.NewEncoder(b).Encode(m)
	fmt.Println("11111", b.Len())

	e := new(bytes.Buffer)
	gob.NewEncoder(e).Encode(d)
	fmt.Println("11111", e.Len())

	go func() {
		http.ListenAndServe("0.0.0.0:9999", nil)
	}()

	select {}
}
