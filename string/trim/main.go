package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaa/cc/vv"
	fmt.Println(strings.TrimRight(s, "/"))

	d := "/aaa/cc/vv"
	fmt.Println(strings.TrimLeft(d, "/"))

	d = "/aaa/cc/vv"
	fmt.Println(strings.TrimLeft(d, "/"))

	e := "/aaa/cc/vv"
	fmt.Println(strings.TrimPrefix(e, "/"))

	e = "aaa/cc/vv"
	fmt.Println(strings.TrimPrefix(e, "/"))
}
