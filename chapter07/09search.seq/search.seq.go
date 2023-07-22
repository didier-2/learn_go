package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	size := 10
	for i := 0; i < 5; i++ {
		arr := generateRandom(size)
		fmt.Println(arr)
	}
}

func generateRandom(size int) []int64 {
	arr := make([]int64, 0, size)
	for i := 0; i < size; i++ {
		i, _ := rand.Int(rand.Reader, big.NewInt(50))
		arr = append(arr, i.Int64())
	}
	return arr
}
