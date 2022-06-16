package main

import "fmt"

func main() {
	fmt.Println(NilStringSlice())
	fmt.Println(len(NilStringSlice()))
}

func NilStringSlice() []string {
	return nil
}
