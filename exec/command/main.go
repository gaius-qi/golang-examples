package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("mv", "foo.log", "logs/foo.log").CombinedOutput()
	if err != nil {
		fmt.Println("move failed: ", err)
	}

	fmt.Println(out)
}
