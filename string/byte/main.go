package main

import (
	"fmt"
)

func main() {
	s := "test"
	b := []byte(s)

	fmt.Println("str to []byte: ", b)
	fmt.Println("[]byte to str: ", string(b))
}
