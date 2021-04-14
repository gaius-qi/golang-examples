package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/avast/retry-go"
)

func main() {
	num := 1

	err := retry.Do(
		func() error {
			if num > 0 {
				fmt.Println(num)
				num += num
				return errors.New("retry")
			}
			return nil
		},
		retry.Delay(1*time.Second),
	)

	if err != nil {
		panic(err)
	}
}
