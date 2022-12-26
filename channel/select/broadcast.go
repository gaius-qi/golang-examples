package main

import (
	"fmt"
	"time"
)

func main() {
	N := 20
	exit := make(chan struct{})
	done := make(chan struct{}, N)

	// start N worker goroutines
	for i := 0; i < N; i++ {
		go func(n int) {
			for {
				select {
				// wait for exit signal
				case <-exit:
					fmt.Printf("worker goroutine #%d exit\n", n)
					done <- struct{}{}
					return
				case <-time.After(time.Second):
					fmt.Printf("worker goroutine #%d is working...\n", n)
				}
			}
		}(i)
	}

	time.Sleep(4 * time.Second)

	// broadcast exit signal
	close(exit)

	// wait for all worker goroutines exit
	for i := 0; i < N; i++ {
		<-done
	}
	fmt.Println("main goroutine exit")
}
