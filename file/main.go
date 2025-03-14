package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.Stat("./test")

	fmt.Println("file:", f.Sys())

	buf := make([]byte, 2)
	for {
		n, err := file.Read(buf)
		fmt.Println("-------", n, err)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}
	}
}
