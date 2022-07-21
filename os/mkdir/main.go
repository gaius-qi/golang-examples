package main

import (
	"fmt"
	"os"
)

func main() {
	err1 := os.MkdirAll("", 0755)
	fmt.Println(err1)

	err2 := os.MkdirAll("main.go", 0755)
	fmt.Println(err2)

	err3 := os.MkdirAll("aa", 0755)
	fmt.Println(err3)
}
