package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := errgroup.Group{}

	g.Go(func() error {
		return errors.New("goroutine 1")
	})

	g.Go(func() error {
		time.Sleep(3 * time.Second)
		return errors.New("goroutine 2")
	})

	werr := g.Wait()
	fmt.Println(werr)

	g2, ctx := errgroup.WithContext(context.Background())
	g2.Go(func() error {
		defer ctx.Done()
		return errors.New("goroutine 1")
	})

	g2.Go(func() error {
		defer ctx.Done()
		time.Sleep(3 * time.Second)
		return errors.New("goroutine 2")
	})

	werr2 := g2.Wait()
	fmt.Println(werr2)
}
