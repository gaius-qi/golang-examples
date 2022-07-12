package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "/download"
	fmt.Println(strings.HasPrefix(a, "/download"))
}
