package main

import (
	"fmt"
	"time"
)

func main() {
	var done chan struct{}

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

loop:
	for {
		select {
		case <-done:
			fmt.Println("chan close")
			break loop
		default:
			fmt.Println("....................")
		}
	}
}
