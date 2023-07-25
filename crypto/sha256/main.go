package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("127.0.0.1"))
	h.Write([]byte("foo"))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
