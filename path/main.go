package main

import (
	"fmt"
	"path"
)

func main() {
	p1 := "//b/c"
	fmt.Println(path.Clean(p1))

	p2 := "b/c"
	fmt.Println(path.Clean(p2))

	p3 := "/b/c.jpg"
	fmt.Println(path.Clean(p3))

	p4 := "b/c.jpg"
	fmt.Println(path.Clean(p4))
}
