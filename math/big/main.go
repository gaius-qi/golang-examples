package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(Cmp(1.0, 1.0))
}

func Cmp(a float64, b float64) bool {
	return big.NewFloat(a).Cmp(big.NewFloat(b)) >= 0
}
