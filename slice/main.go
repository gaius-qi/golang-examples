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
	foods[1] = "apple"
	fmt.Println(foods)

	// var person []string
	// person[10] = "lucy"
	// fmt.Println(len(person))

	// rooms := make([]int, 0, 2)
	// things := []int{0, 1, 2, 3}
	// for _, thing := range things {
	// fmt.Println("place thing", thing)
	// rooms = append(rooms, thing)
	// }
	// fmt.Println("rooms", rooms)

	// box := []int{}
	// items := []int{0, 1, 2, 3}
	// for index, item := range items {
	// if index > 2 {
	// break
	// }

	// box = append(box, item)
	// }
	// fmt.Println("box", box)

	numbers := []int{1}
	fmt.Println("numbers: ", numbers[1:])
}
