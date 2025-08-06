package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Node struct {
	IP                 string  `json:"ip,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	CachedLayers       []Layer `json:"cachedLayers"`
	SchedulerClusterID int     `json:"schedulerClusterId"`
}

type Layer struct {
	URL string `json:"url"`
}

type Image struct {
	Nodes []Node `json:"peers"`
}

func main() {
	// Read the JSON file
	data, err := ioutil.ReadFile("jsonformatter.txt")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Parse JSON data
	var image Image
	err = json.Unmarshal(data, &image)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Count layers for each node
	for i, node := range image.Nodes {
		var identifier string
		if node.IP != "" && node.Hostname != "" {
			identifier = fmt.Sprintf("Node %d (IP: %s, Hostname: %s)", i+1, node.IP, node.Hostname)
		} else {
			identifier = fmt.Sprintf("Node %d", i+1)
		}
		fmt.Printf("%s has %d layers\n", identifier, len(node.CachedLayers))
	}
}
