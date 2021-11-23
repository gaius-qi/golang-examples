package main

import (
	"fmt"

	"github.com/Showmax/go-fqdn"
)

func main() {
	fqdn, err := fqdn.FqdnHostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(fqdn)
}
