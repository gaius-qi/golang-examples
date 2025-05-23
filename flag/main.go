package main

import (
	"flag"
	"fmt"
)

var (
	name  = flag.String("name", "default", "name of the user")
	debug = flag.Bool("debug", false, "print debug data")
)

func main() {

	flag.Parse()
	if len(flag.Args()) < 1 {
		panic("Usage")
	}

	fmt.Println("111111111", *debug)
	fmt.Println("111111111", *name)
}
