package main

import "fmt"

type Food interface{}

func main() {
	var food Food

	fmt.Println(food == nil)
}
