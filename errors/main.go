package main

import (
	"errors"
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

func main() {
	a := errors.New("a")
	e1 := pkgerrors.Wrap(a, "aaa")
	e2 := fmt.Errorf("aaa: %w", a)

	fmt.Println(e1)
	fmt.Println(e2)

	e3 := pkgerrors.Wrap(e1, "bbb")
	e4 := fmt.Errorf("bbb: %w", e2)

	fmt.Println(e3)
	fmt.Println(e4)

	fmt.Println(errors.Is(e3, a))
	fmt.Println(errors.Is(e4, a))
}
