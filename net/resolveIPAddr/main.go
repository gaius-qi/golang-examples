package main

import (
	"fmt"
	"net"
)

func main() {
	r1, err := net.ResolveIPAddr("ip4:icmp", "39.156.66.10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ipv4: %s\n", r1.String())

	r2, err := net.ResolveIPAddr("ip6:icmp", "2402:4e00::b")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ipv6: %s\n", r2.String())
}
