package main

import "fmt"

func main() {
	str := "abcd"
	for i, v := range str {
		fmt.Println(i, v, str[i])
	}
}
