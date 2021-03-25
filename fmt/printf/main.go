package main

import "fmt"

type student struct {
	id   int32
	name string
}

func main() {
	a := &student{id: 1, name: "jack"}

	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%#v\n", a)
}
