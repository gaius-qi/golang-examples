package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"syscall"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   time.Second * 20,
		ResponseHeaderTimeout: time.Second * 30,
		IdleConnTimeout:       time.Second * 300,
		MaxIdleConnsPerHost:   500,
		ReadBufferSize:        32 << 10,
		WriteBufferSize:       32 << 10,
		DisableCompression:    true,
		TLSClientConfig:       &tls.Config{},
	},
	Timeout: time.Hour,
}

func main() {
	proxyURL, _ := url.Parse("http://127.0.0.1:65001")
	httpClient.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)

	req, _ := http.NewRequest("GET", "https://www.google.com", nil)

	resp, err := httpClient.Do(req)
	if err != nil {
		if errors.Is(err, syscall.ECONNREFUSED) {
			fmt.Println("Connection refused")
		}
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	fmt.Println(resp.Status)
}
