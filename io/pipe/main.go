package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
	r, w := io.Pipe()

	go func() {
		for {
			time.Sleep(time.Second)
			w.Write([]byte(time.Now().String()))
		}
	}()

	for {
		dataRead := make([]byte, 256)
		n, _ := r.Read(dataRead)
		fmt.Println(string(dataRead[:n]))
	}
}
