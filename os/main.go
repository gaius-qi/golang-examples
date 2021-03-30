package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./user/cache")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("create file success: %v", f)
}
