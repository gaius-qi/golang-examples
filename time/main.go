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
	diff := at2.Sub(at1)
	fmt.Println(diff.Nanoseconds())

	t3, err := time.Parse(time.RFC3339Nano, fmt.Sprint(time.Now().Nanosecond()))
	fmt.Println("-----", t3, err)
}
