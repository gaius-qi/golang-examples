package main

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type car struct{}

func main() {
	required := mapset.NewSet[car]()
	required.Add(car{})
}
