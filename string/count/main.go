package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "aa|bb|cc"
	str2 := "|"
	fmt.Println(strings.Count(str1, str2))
}
