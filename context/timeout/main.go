package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	asyncCall()
}

func asyncCall() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go func(ctx context.Context) {
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
