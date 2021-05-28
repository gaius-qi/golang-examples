package main

import "time"

func main() {
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-tick.C:
			return
		}
	}
}
