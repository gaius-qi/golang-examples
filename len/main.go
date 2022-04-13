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

	ns2 := make([]*User, 0, 3)
	fmt.Println(len(ns2))

	ns2 = append(ns2, &User{Name: "jack"}, &User{Name: "lucy"})
	fmt.Println(ns2, ns2[0].Name, ns2[1].Name, len(ns2), cap(ns2))

	ns2 = ns2[:0]
	fmt.Println(ns2, len(ns2), cap(ns2))

	ns2 = append(ns2, &User{Name: "gaius"})
	fmt.Println(ns2, ns2[0].Name, len(ns2), cap(ns2))

	ns2 = append(ns2, &User{Name: "gaius"})
	ns2 = append(ns2, &User{Name: "gaius"})
	ns2 = append(ns2, &User{Name: "gaius"})
	ns2 = append(ns2, &User{Name: "gaius"})
	fmt.Println(ns2, len(ns2), cap(ns2))
}
