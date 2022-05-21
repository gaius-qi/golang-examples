package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	cert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	conf := &tls.Config{
		RootCAs: certPool,
	}

	server := &http.Server{
		Addr:      ":8080",
		Handler:   &Proxy{},
		TLSConfig: conf,
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
