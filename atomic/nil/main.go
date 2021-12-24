package main

import (
	"fmt"

	"go.uber.org/atomic"
)

type food interface{}

func main() {
	av := atomic.Value{}
	set(av, nil)

	rawFood := av.Load()
	fmt.Println("rawFood: ", rawFood)

	car := atomic.Value{}
	rawCar := car.Load()
	if rawCar == nil {
		fmt.Println("rawCar", rawCar)
	}
}

func set(av atomic.Value, item food) {
	av.Store(item)
}
