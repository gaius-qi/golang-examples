package main

import (
	"fmt"
	"time"
)

func main() {
	times := 10000000
	now := time.Now()

	for i := 0; i < times; i++ {
		if i == 100 {
			fmt.Println("=============")
		}
	}

	fmt.Printf("times %d, cost %dms\n", times, time.Since(now).Milliseconds())
}
