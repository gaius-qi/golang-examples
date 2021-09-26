package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	mutliGoroutineTimeout()
}

func asyncCall() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		time.Sleep(3 * time.Second)
		fmt.Println("reuqest!!!")
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("call successfully!!!")
		return
	case <-time.After(2 * time.Second):
		fmt.Println("timeout!!!")
		return
	}
}

func mutliGoroutineTimeout() {
	run()
	time.Sleep(10 * time.Second)
}

func run() {
	done := make(chan struct{})

	go func() {
		defer close(done)
		time.Sleep(1 * time.Second)
		fmt.Println("running")
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
		return
	case <-done:
		fmt.Println("done")
		return
	}
}
