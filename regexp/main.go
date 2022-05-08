package main

import (
	"fmt"
	"regexp"
)

func main() {
	regx1, _ := regexp.Compile("file-server*")
	r1 := regx1.MatchString("http://file-server.dragonfly-e2e.svc/kind/etc/containerd/config.toml")
	fmt.Println("result:", r1)

	regx2, _ := regexp.Compile("file")
	r2 := regx2.MatchString("http://file-server.dragonfly-e2e.svc/kind/etc/containerd/config.toml")
	fmt.Println("result:", r2)
}
