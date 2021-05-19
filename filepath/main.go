package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base("/foo/bar/baz.js"))
	fmt.Println(filepath.Dir("/foo/bar/baz.js"))
	fmt.Println(filepath.Join("/foo", "bar/baz.js"))
}
