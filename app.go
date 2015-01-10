package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	max := *big.NewInt(1000)
	fmt.Println(rand.Int(rand.Reader, &max))
}
