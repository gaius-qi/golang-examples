package main

import (
	"fmt"
	"net"
)

func main() {
	host1, port1, _ := net.SplitHostPort("0.0.0.0:1010")
	fmt.Println(host1, port1)

	host2, port2, _ := net.SplitHostPort("localhost:1010")
	fmt.Println(host2, port2)
}