package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaa/cc/vv"
	fmt.Println(strings.TrimRight(s, "/"))
}
