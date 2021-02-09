package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
)

func main() {
	// cert.Certificate is a chain of one or more certificates, leaf first.
	cert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	if err != nil {
		panic(err)
	}

	leaf, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Leaf %+v\n", leaf.IsCA)
}
