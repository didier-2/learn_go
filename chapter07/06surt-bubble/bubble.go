package main

import (
	"fmt"
	"math/rand"
)

func bubble(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
		fmt.Println("中间状态", *arr)
	}
	fmt.Println("最终状态", *arr)
}
func main() {
	arrSize := 10
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}
	bubble(&arr)
}
