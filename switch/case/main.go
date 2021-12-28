package main

import "fmt"

func main() {
	code := 2
	switch code {
	case 1:
		fmt.Println("1")
		return
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println("default")
	}

	fmt.Println("last line")
}
