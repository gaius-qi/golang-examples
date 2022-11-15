package main

import "time"

func main() {
	for i := 0; i < 12; i++ {
		go func() {
			for {
			}
		}()
	}

	time.Sleep(time.Hour)
}
