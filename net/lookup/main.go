package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	addrs, err := net.LookupHost(host)
	if err != nil {
		panic(err)
	}
	fmt.Println(addrs)

	for _, addr := range addrs {
		hosts, err := net.LookupAddr(addr)
		if err != nil || len(hosts) == 0 {
			continue
		}
		fmt.Println(strings.TrimSuffix(hosts[0], "."))
	}

	var hosts []string
	hosts, err = net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("hosts: %#v\n", hosts)

	var ips []net.IP
	ips, err = net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ips: %s\n", ips)
}
