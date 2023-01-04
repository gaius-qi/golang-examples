package main

import (
	"fmt"
	"net"
	"os"
	"time"

	fastping "github.com/tatsushid/go-fastping"
)

func main() {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip6:icmp", "2402:4e00::b")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}

	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
