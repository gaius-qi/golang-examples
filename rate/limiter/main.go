package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limit := rate.NewLimiter(3, 5)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for i := 0; ; i++ {
		fmt.Printf("%03d %s\n", i, time.Now().Format("2008-01-01 15:04:04.000"))
		err := limit.Wait(ctx)
		if err != nil {
			fmt.Println("err: ", err.Error())
			return
		}
	}
}
