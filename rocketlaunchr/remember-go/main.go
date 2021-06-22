package main

import (
	"fmt"

	"github.com/rocketlaunchr/remember-go"
)

type Key struct {
	Search string
	Page   int `json:"page"`
}

func main() {
	key := remember.CreateKeyStruct(Key{"golang", 2})

	fmt.Println(key)
}
