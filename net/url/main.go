package main

import (
	"fmt"
	"net/url"
)

func main() {
	if _, err := url.ParseRequestURI("http://google.com"); err != nil {
		fmt.Println("1", err)
	}

	if _, err := url.ParseRequestURI("google.com"); err != nil {
		fmt.Println("2", err)
	}

	if _, err := url.ParseRequestURI("/foo/bar"); err != nil {
		fmt.Println("3", err)
	}

	if _, err := url.ParseRequestURI("http://"); err != nil {
		fmt.Println("4", err)
	}

	if _, err := url.ParseRequestURI("http://google.com/ss/cdsa?cdsa=cd"); err != nil {
		fmt.Println("5", err)
	}
}
