package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	targetURL := url.URL{
		Scheme: "https",
		Host:   "google",
		Path:   "api/v1",
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, targetURL.String(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(targetURL.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// The HTTP 206 Partial Content success status response code indicates that
	// the request has succeeded and the body contains the requested ranges of data, as described in the Range header of the request.
	// Refer to https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/206.
	if resp.StatusCode/100 != 2 {
		panic(resp.StatusCode)
	}

	fmt.Println(io.ReadAll(resp.Body))
}
