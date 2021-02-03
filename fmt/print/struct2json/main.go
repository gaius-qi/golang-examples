package main

import (
	"encoding/json"
	"fmt"
)

type Apple struct {
	Color  string
	Weight int
}

func main() {
	a := &Apple{
		Color:  "red",
		Weight: 10,
	}

	s, _ := json.MarshalIndent(a, "", "\t")
	fmt.Printf("%s", s)
}
