package main

import (
	"fmt"
	"reflect"
)

const (
	CategoryBooks  = iota // 0
	CategoryHealth        // 1
	_
	CategoryClothing // 3
)

type Stereotype int

const (
	TypicalNoob           Stereotype = iota + 1 // 1
	TypicalHipster                              // 2
	TypicalUnixWizard                           // 3
	TypicalStartupFounder                       // 4
)

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

const (
	FooWeight float64 = 0.1
	BarWeight         = 0.2
)

func main() {
	fmt.Println(CategoryBooks, CategoryHealth, CategoryClothing)
	fmt.Println(TypicalNoob, TypicalHipster, TypicalUnixWizard, TypicalStartupFounder)
	fmt.Println(IgEggs, IgChocolate, IgNuts, IgStrawberries, IgShellfish)
	fmt.Println(reflect.TypeOf(FooWeight), reflect.TypeOf(BarWeight))
}
