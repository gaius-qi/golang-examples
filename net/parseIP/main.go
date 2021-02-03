package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.ParseIP("0.0.0.0"))
	fmt.Println(net.IPv4zero.Equal(net.ParseIP("0.0.0.0")))
}
