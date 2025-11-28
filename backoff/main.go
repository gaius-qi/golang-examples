package main

import (
	"github.com/cenkalti/backoff/v5"
)

func main() {
	ticker := backoff.NewTicker(backoff.NewExponentialBackOff())
}
