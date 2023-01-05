package main

import (
	"fmt"
	"net"
	"time"

	fastping "github.com/tatsushid/go-fastping"
)

func main() {
	p := fastping.NewPinger()
	// r0, err := net.ResolveIPAddr("ip6:icmp", "2402:4e00::b")
	r0, err := net.ResolveIPAddr("ip4:icmp", "39.156.66.10")
	if err != nil {
		panic(err)
	}

	p.MaxRTT = 2 * time.Second
	p.Debug = true
	p.AddIPAddr(r0)
	if _, err := p.Network("udp"); err != nil {
		panic(err)
	}

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}

	p.OnIdle = func() {
		fmt.Println("finish")
	}

	if err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
