package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	config := NewConfig()
	err := config.LoadYaml("./config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(config, config.Name, config.Nodes[0].ID)
}

const (
	defaultName = "default"
)

type Config struct {
	Name  string    `yaml:"name,omitempty" json:"name,omitempty"`
	Nodes []*Weight `yaml:"nodes,omitempty" json:"nodes,omitempty"`
}

type Weight struct {
	ID     string
	Weight int
}

func NewConfig() *Config {
	// don't set Supernodes as default value, the SupernodeLocator will
	// do this in a better way.
	return &Config{
		Name: defaultName,
	}
}

// Load loads properties from config file.
func (c *Config) LoadYaml(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to load yaml %s when reading file: %v", path, err)
	}
	if err = yaml.Unmarshal(content, c); err != nil {
		return fmt.Errorf("failed to load yaml %s: %v", path, err)
	}
	return nil
}
