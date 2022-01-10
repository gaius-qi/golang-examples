package main

import (
	"encoding/json"
	"fmt"
)

type config struct {
	LoadLimit   uint32 `yaml:"loadLimit" mapstructure:"loadLimit" json:"load_limit" binding:"omitempty,gte=1,lte=5000"`
	NetTopology string `yaml:"netTopology" mapstructure:"netTopology" json:"net_topology"`
}

func main() {
	b, err := json.Marshal(config{
		NetTopology: "foo",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(b)
}
