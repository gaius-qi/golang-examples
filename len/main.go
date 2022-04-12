package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var n []*User
	fmt.Println(len(n))

	ns1 := make([]*User, 2)
	fmt.Println(len(ns1))

	ns2 := make([]*User, 0, 1)
	fmt.Println(len(ns2))
}
