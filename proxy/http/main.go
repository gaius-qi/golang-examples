package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	n := 50
	target := "http://127.0.0.1:8081/medium"
	proxy := "http://127.0.0.1:4001"

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(target string) {
			defer wg.Done()
			err := download(target, proxy)
			if err != nil {
				fmt.Printf("Failed to download %s: %v\n", target, err)
			}
		}(target)
	}

	wg.Wait()
	fmt.Println("All downloads are done")
}

func download(target string, proxy string) error {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
		Timeout: time.Hour,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	if _, err := io.ReadAll(resp.Body); err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return err
	}

	return nil
}
