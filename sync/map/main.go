package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}

	// m.Store("foo", "bar1")
	// v1, _ := m.Load("foo")
	// fmt.Println("v1:", v1)
	// v2, _ := m.LoadOrStore("foo", "bar2")
	// fmt.Println("v2:", v2)
	// m.Store("foo", "bar3")
	// v3, _ := m.Load("foo")
	// fmt.Println("v3:", v3)

	for i := 0; i < 20; i++ {
		m.Store(i, i)
	}

	v1 := []interface{}{}
	m.Range(func(k, v interface{}) bool {
		v1 = append(v1, v)
		return true
	})
	fmt.Println("-------------")
	fmt.Println(v1)
	fmt.Println("-------------")

	v2 := []interface{}{}
	m.Range(func(k, v interface{}) bool {
		v2 = append(v2, v)
		return true
	})
	fmt.Println("-------------")
	fmt.Println(v2)
	fmt.Println("-------------")

	v3 := []interface{}{}
	m.Range(func(k, v interface{}) bool {
		v3 = append(v3, v)
		return true
	})
	fmt.Println("-------------")
	fmt.Println(v3)
	fmt.Println("-------------")
}
