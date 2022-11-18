package main

import (
	"fmt"
	"reflect"
)

type car struct {
	name string
}

func main() {
	foo := []*car{
		{
			name: "foo",
		},
	}

	bar := []*car{
		{
			name: "bar",
		},
	}

	if reflect.DeepEqual(foo, bar) {
		fmt.Println("equal")
		return
	}

	fmt.Println("not equal")

	a := []string{"a", "c"}
	b := []string{"a", "b"}

	if reflect.DeepEqual(a, b) {
		fmt.Println("equal")
		return
	}

	fmt.Println("not equal")
}
