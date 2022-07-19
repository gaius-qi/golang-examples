package main

import "fmt"

func main() {
	s1 := []string{"a"}

	fmt.Println(len(s1) == 1)
	fmt.Println(s1[0])
	fmt.Println(s1[1:])

	s2 := []string{"a", "b", "c", "d"}
	fmt.Println(s2[:1])
}
