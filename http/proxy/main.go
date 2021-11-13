package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: &Proxy{},
	}
	server.ListenAndServe()
}

type Proxy struct {
}

func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Method: %s \n", r.Method)
	fmt.Printf("r.URL.String(): %s \n", r.URL.String())
	fmt.Printf("r.URL.HOST: %s \n", r.URL.Host)
	fmt.Printf("r.HOST: %s \n", r.Host)
	fmt.Println("=======================")
}
