package main

import (
	"time"

	"github.com/go-echarts/statsview"
)

func main() {
	mgr := statsview.New()

	// Start() runs a HTTP server at `localhost:18066` by default.
	go mgr.Start()

	// Stop() will shutdown the http server gracefully
	// mgr.Stop()

	// busy working....
	time.Sleep(10 * time.Minute)
}
