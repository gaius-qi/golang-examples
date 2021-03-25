package main

import (
	"fmt"
	"strconv"
)

type Animaler interface {
	eat()
}

type Animal struct {
	name string
	age  int
	food string
}

func (animal Animal) eat() {
	fmt.Println("Name: " + animal.name + ", Age: " + strconv.Itoa(animal.age) + ", Eat: " + animal.food)
}

type Cat struct {
	Animal
	time int
}

func (cat Cat) sleep() {
	fmt.Println("Name: " + cat.name + ", Sleep: " + strconv.Itoa(cat.time))
}

type Dog struct {
	Animal
	plays string
}

func (dog Dog) play() {
	fmt.Println("Name: " + dog.name + ", Play:" + dog.plays)
}

func testInterface(animaler Animaler) {
	animaler.eat()
}

func main() {
	cat := Cat{Animal: Animal{name: "Jack", age: 2, food: "fish"}, time: 8}
	cat.sleep()
	testInterface(cat)

	dog := Dog{Animal: Animal{name: "Lucy", age: 2, food: "meat"}, plays: "football"}
	dog.play()
	testInterface(dog)
}
