package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	host1, port1, _ := net.SplitHostPort("0.0.0.0:1010")
	fmt.Println(host1, port1)

	host2, port2, _ := net.SplitHostPort("localhost:1010")
	fmt.Println(host2, port2)

	host3, port3, err := net.SplitHostPort(":1010")
	if err != nil {
		panic(err)
	}
	fmt.Println(host3, port3)

	u, err := url.Parse("vsock://0.0.0.0:1010")
	if err != nil {
		panic(err)
	}

	host4, port4, err := net.SplitHostPort(u.Host)
	fmt.Println(host4, port4, err)
}
