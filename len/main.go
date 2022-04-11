package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var n []*User
	fmt.Println(len(n))
}
