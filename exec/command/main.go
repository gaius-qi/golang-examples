package main

import (
	"fmt"
	"os/exec"
)

func main() {
	o1, err1 := exec.Command("mv", "foo.log", "logs/foo.log").CombinedOutput()
	if err1 != nil {
		fmt.Println("move failed: ", err1)
	}

	fmt.Println(o1)

	o2, err2 := exec.Command("kubectl", "-n", "dragonfly-system", "exec", "-c", "seed-peer", "dragonfly-seed-peer-1", "--", "sh", "-c", "sha256sum /var/lib/dragonfly/581e9ef4c4e8b3363ca37521f451f04f924fc2f1fd1834143e4d490396f389c2/*/data").CombinedOutput()
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(string(o2))
}
