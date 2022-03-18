package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("", 0755)
	fmt.Println(err)
}
