package main

import "fmt"

func main() {
	cars := make([]string, 0, 2)
	cars = append(cars, "a")
	fmt.Println(cars)
	cars = append(cars, "b")
	fmt.Println(cars)
	cars = append(cars, "c")
	fmt.Println(cars)
	cars = append(cars, "d")
	fmt.Println(cars)

	foods := (make([]string, 2))
	foods[3] = "apple"
	fmt.Println(foods)
}
