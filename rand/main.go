package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))

	permutation := rand.Perm(10)[:10]
	fmt.Println(permutation)

	permutation2 := rand.Perm(2)[:2]
	fmt.Println(permutation2)

	permutation3 := rand.Perm(10)[:2]
	fmt.Println(permutation3)
}
