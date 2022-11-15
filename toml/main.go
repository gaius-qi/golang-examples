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

type MyDoc struct {
	Mode ModeType
}

func main() {
	var doc MyDoc
	data := `mode = "aa"`
	if err := toml.Unmarshal([]byte(data), &doc); err != nil {
		panic(err)
	}

	fmt.Printf("%v", doc.Mode)
}
