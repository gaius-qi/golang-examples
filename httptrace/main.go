package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptrace"
	"os"
	"strings"
	"time"
)

func main() {

	dialer := &net.Dialer{Timeout: 7 * time.Second}
	// For security's sake, sorry, I can't provide detailed domain names
	url := os.Args[1]
	// json content
	payload := strings.NewReader("")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	trace := &httptrace.ClientTrace{
		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			fmt.Printf("dns start: %v\n", time.Now())
		},
		DNSDone: func(dnsDoneInfo httptrace.DNSDoneInfo) {
			fmt.Printf("dns end: %v\n", time.Now())
		},
		WroteHeaders: func() {
			fmt.Printf("wrote headers: %v\n", time.Now())
		},
		WroteRequest: func(wroteRequestInfo httptrace.WroteRequestInfo) {
			fmt.Printf("wrote request info: %v\n", time.Now())
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Printf("connect done: %v\n", time.Now())
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	transport := &http.Transport{

		DialContext:         (dialer).DialContext,
		Dial:                (dialer).Dial,
		TLSHandshakeTimeout: 2 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(10000) * time.Millisecond,
	}

	now := time.Now()

	res, err := client.Do(req)
	res.Body.Close()

	cost := time.Since(now)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("client cost: %v\n", cost.Seconds()*1e3)
}
