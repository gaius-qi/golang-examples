package main

import (
	"fmt"
	"time"
)

const (
	TaskStatusWaiting int = iota
	TaskStatusRunning
	TaskStatusSeeding
	TaskStatusSuccess
	TaskStatusZombie
	TaskStatusFail
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			switch TaskStatusWaiting {
			case TaskStatusRunning, TaskStatusSeeding:
			case TaskStatusSuccess:
				fmt.Println("success")
				return
			case TaskStatusWaiting:
				if false {
					fmt.Println("waiting true")
					break
				}

				fallthrough
			default:
				fmt.Println("default")
				return
			}
		}
	}
}
