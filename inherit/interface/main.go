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

// 实现 eat() 方法
func (animal Animal) eat() {
	fmt.Println(animal.name + "今年" + strconv.Itoa(animal.age) + "岁了，" + "喜欢吃" + animal.food)
}

type Cat struct {
	Animal
	time int
}

type Dog struct {
	Animal
	plays string
}

// 猫独有的方法
func (cat Cat) sleep() {
	fmt.Println("我叫" + cat.name + "， 我能睡" + strconv.Itoa(cat.time) + "分钟")
}

// 狗独有的方法
func (dog Dog) play() {
	fmt.Println("我叫" + dog.name + "我喜欢玩" + dog.plays)
}

func testInterface(animaler Animaler) {
	animaler.eat()
}

func main() {
	cat := Cat{Animal: Animal{name: "咪咪", age: 2, food: "鱼"}, time: 8}
	cat.sleep()
	testInterface(cat)

	dog := Dog{Animal: Animal{name: "大黄", age: 2, food: "骨头"}, plays: "球球"}
	dog.play()
	testInterface(dog)
}
