package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./test/cache")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("create file success: %v", f)

	dir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("user cache dir: %v", dir)
}
