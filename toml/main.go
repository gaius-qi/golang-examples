package main

import (
	"fmt"

	"github.com/pelletier/go-toml/v2"
)

type ModeType string

func (m *ModeType) UnmarshalText(b []byte) error {
	*m = ModeType(b)
	return nil
}

func (m *ModeType) MarshalText() ([]byte, error) {
	return []byte(string(*m)), nil
}

type MyDoc struct {
	Mode ModeType
}

func main() {
	var doc MyDoc
	data := `mode = "aa"`
	if err := toml.Unmarshal([]byte(data), &doc); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", doc.Mode)

	b, err := toml.Marshal(doc)
	if err != nil {
		panic(err)
	}

	fmt.Println("byte: ", b)

	var doc2 MyDoc
	if err := toml.Unmarshal(b, &doc2); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", doc.Mode)
}
