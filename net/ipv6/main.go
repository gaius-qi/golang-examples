package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.IPv6loopback.String())
}
