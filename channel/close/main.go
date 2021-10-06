package main

import "fmt"

func main() {
	c := make(chan string, 1)
	close(c)
	fmt.Println(c)

	c <- "a"
	fmt.Println(c)
}
