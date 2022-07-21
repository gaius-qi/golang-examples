package main

import (
	"fmt"

	pkglog "d7y.io/dragonfly/v2/pkg/log"
)

func main() {
	err := pkglog.SetupDaemon("./test", false, false)
	fmt.Println(err)
}
