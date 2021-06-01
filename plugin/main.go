package main

import (
	"fmt"
	"os"
	"plugin"
)

type Worker interface {
	Run()
}

func init() {
	fmt.Println("init main...")
}

func main() {
	fmt.Println("main start...")

	fmt.Println("plugin open...")
	plug, err := plugin.Open("./plugins/plugin_worker.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("plugin lookup...")
	doc, err := plug.Lookup("Worker")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	worker, ok := doc.(Worker)
	if !ok {
		fmt.Println("unexpected type from module")
		os.Exit(1)
	}

	worker.Run()
}
