package main

import (
	"errors"
	"fmt"

	"stathat.com/c/consistent"
)

func main() {
	hashring := consistent.New()

	hashring.Add("scheduler-0")
	hashring.Add("scheduler-1")
	hashring.Add("scheduler-2")
	hashring.Add("scheduler-3")
	hashring.Add("scheduler-4")
	hashring.Add("scheduler-5")
	hashring.Add("scheduler-6")

	fmt.Printf("Members: %v\n", hashring.Members())

	s, err := hashring.Get("task-0")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get key {task-0}: %s\n", s)

	s, err = hashring.Get("task-1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get key {task-1}: %s\n", s)

	s, err = hashring.Get("task-2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get key {task-2}: %s\n", s)

	// circle, err := GenerateCircle(hashring)
	// if err != nil {
	// panic(err)
	// }

	// fmt.Printf("circle: %#v\n", circle)
}

func GenerateCircle(hashring *consistent.Consistent) (map[string]string, error) {
	members := hashring.Members()
	circle := make(map[string]string, len(members))

	for i := 0; i <= len(members)*10; i++ {
		key := fmt.Sprint(i)
		member, err := hashring.Get(key)
		if err != nil {
			fmt.Printf("hashring get member failed: %s", err.Error())
			continue
		}

		if _, ok := circle[member]; !ok {
			circle[member] = key
			if len(circle) == len(members) {
				return circle, nil
			}
		}
	}

	return nil, errors.New("can not generate circle")
}
