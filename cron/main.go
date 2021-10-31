package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	i := 0
	c := cron.New()
	spec := fmt.Sprintf("@every %s", 3*time.Second)

	if _, err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	}); err != nil {
		fmt.Println(err)
		return
	}

	c.Start()

	select {}
}
