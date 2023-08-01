package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var totolCompare int = 0

func main() {
	size := 1000
	arr := generateRandom(size)
	fmt.Println(arr)
	//1、501在不在
	//2、888在不在
	//3、900在不在
	//4、3在不在
	fmt.Println(search(&arr, 501))
	fmt.Println(search(&arr, 888))
	fmt.Println(search(&arr, 900))
	fmt.Println(search(&arr, 3))
	fmt.Println("总比较次数：", totolCompare)

}

func search(arrp *[]int64, targetNum int64) bool {
	for _, v := range *arrp {
		if v == targetNum {
			return true
		}
	}
	return false
}

func generateRandom(size int) []int64 {
	arr := make([]int64, 0, size)
	for i := 0; i < size; i++ {
		i, _ := rand.Int(rand.Reader, big.NewInt(3000))
		arr = append(arr, i.Int64())
	}
	return arr
}
