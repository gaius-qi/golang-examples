package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		c <- "a"
	}()

	select {
	case v := <-c:
		fmt.Println(v)
	}
}
