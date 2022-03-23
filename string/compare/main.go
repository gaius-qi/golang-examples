package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("A10", "a10"), strings.EqualFold("A10", "A11"))
}
