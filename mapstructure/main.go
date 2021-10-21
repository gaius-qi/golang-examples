package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type BIO struct {
	ID uint
}

func main() {
	var config map[string]interface{}
	mapstructure.Decode(&struct {
		Name string
		BIO  BIO
	}{
		Name: "foo",
		BIO: BIO{
			ID: 1,
		},
	}, &config)

	fmt.Printf("%+v\n", config)
}
