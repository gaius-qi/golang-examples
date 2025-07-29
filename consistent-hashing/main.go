package main

import (
	"errors"
	"fmt"

	"stathat.com/c/consistent"
)

func main() {
	hashring := consistent.New()

	hashring.Add("30.46.243.252:scheduler-1")
	hashring.Add("11.46.101.47:scheduler-2")
	hashring.Add("11.46.129.55:scheduler-3")
	// hashring.Add("21.26.329.11:scheduler-4")
	// hashring.Add("37.11.29.34:scheduler-a")
	// hashring.Add("37.11.323.34:scheduler-b")

	s, err := hashring.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get test: %s\n", s)

	s, err = hashring.Get("test-retry-1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get test-1: %s\n", s)

	s, err = hashring.Get("test-retry-2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get test-2: %s\n", s)

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
