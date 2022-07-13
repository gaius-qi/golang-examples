package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	s := "aaa/cc/vv"
	fmt.Println(strings.TrimRight(s, string(os.PathSeparator)))

	d := "/aaa/cc/vv"
	fmt.Println(strings.TrimLeft(d, string(os.PathSeparator)))

	d = "/aaa/cc/vv"
	fmt.Println(strings.TrimLeft(d, string(os.PathSeparator)))

	e := "/aaa/cc/vv"
	fmt.Println(strings.TrimPrefix(e, string(os.PathSeparator)))

	e = "aaa/cc/vv"
	fmt.Println(strings.TrimPrefix(e, string(os.PathSeparator)))

	fmt.Println(filepath.Join("/", "/cdas/b"))
	fmt.Println(filepath.Join("/", "cdas/b"))
}
