package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(60 * time.Second)

	t1 := time.Now().UnixNano()
	at1 := time.Unix(0, t1)
	time.Sleep(2 * time.Second)
	t2 := time.Now().UnixNano()
	at2 := time.Unix(0, t2)
	aa := at2.Sub(at1)
	fmt.Println(aa.Nanoseconds())
}
