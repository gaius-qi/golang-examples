package main

import (
	"fmt"
	"net"
)

func main() {
	ra, err := net.ResolveIPAddr("ip6:icmp", "2402:4e00::b")
	if err != nil {
		panic(err)
	}

	fmt.Printf("ip: %s\n", ra.String())
}
