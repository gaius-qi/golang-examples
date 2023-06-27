package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-tick.C:
			fmt.Println("tick")
			return
		}
	}
}
