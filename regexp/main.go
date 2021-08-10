package main

import (
	"fmt"
	"regexp"
)

func main() {
	regx, _ := regexp.Compile("file-server*")
	r := regx.MatchString("http://file-server.dragonfly-e2e.svc/kind/etc/containerd/config.toml")
	fmt.Println("result:", r)
}
