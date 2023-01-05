package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	pinger, err := ping.NewPinger("110.242.68.66")
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(false)
	pinger.Count = 1
	pinger.Timeout = time.Second * 5
	if err := pinger.Run(); err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	fmt.Println(stats.AvgRtt, stats.PacketLoss)
}
