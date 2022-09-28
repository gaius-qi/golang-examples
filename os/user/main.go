package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(currentUser)
	fmt.Println(currentUser.Username)
}
