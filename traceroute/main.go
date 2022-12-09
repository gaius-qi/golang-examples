package main

import (
	"fmt"

	"github.com/aeden/traceroute"
)

func main() {
	result, err := traceroute.Traceroute("google.com", &traceroute.TracerouteOptions{})
	if err != nil {
		panic(err)
	}

	for _, hop := range result.Hops {
		fmt.Printf("Number: %d Success: %t Address: %s Host: %s ElapsedTime: %d TTL: %d\n", hop.N, hop.Success, hop.Address, hop.Host, hop.ElapsedTime.Milliseconds(), hop.TTL)
	}
}
