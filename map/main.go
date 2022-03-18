package main

import "fmt"

func main() {
	m := map[string]string{"a": ""}
	v1, ok1 := m["a"]
	fmt.Println(v1, ok1)
	v2, ok2 := m["b"]
	fmt.Println(v2, ok2)
	fmt.Println(v2 == "")
}
