package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	um1, _ := strconv.ParseInt(strconv.Itoa(777), 8, 0)
	fmt.Println(os.FileMode(um1), um1)

	fmt.Println(os.FileMode(0600))

	um2, _ := strconv.ParseInt(strconv.Itoa(300), 8, 0)
	fmt.Println(os.FileMode(um2), um2)

	fmt.Println(os.FileMode(0300))
}
