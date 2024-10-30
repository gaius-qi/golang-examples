package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("https://example.com"))
	h.Write([]byte("foo"))
	h.Write([]byte("bar"))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
