package main

import (
	"fmt"

	"github.com/google/uuid"
	"stathat.com/c/consistent"
)

func main() {
	hashring := consistent.New()

	hashring.Add("30.46.243.252:scheduler-1")
	hashring.Add("11.46.101.47:scheduler-2")
	hashring.Add("11.46.129.55:scheduler-3")
	hashring.Add("21.26.329.11:scheduler-4")
	hashring.Add("37.11.29.34:scheduler-a")
	hashring.Add("37.11.323.34:scheduler-b")

	members := hashring.Members()
	circle := make(map[string]string, len(members))

	for i := 0; i <= len(members)*10; i++ {
		key := uuid.New().String()
		member, err := hashring.Get(key)
		if err != nil {
			panic(err)
		}

		if _, ok := circle[member]; !ok {
			fmt.Printf("%d append memeber %s key %s\n", i, member, key)
			circle[member] = key
			if len(circle) == len(members) {
				break
			}
		}

		fmt.Printf("%d ignore memeber %s key %s\n", i, member, key)
	}

	fmt.Printf("circle: %#v\n", circle)
}
