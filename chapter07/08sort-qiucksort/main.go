package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func quickSort(arr *[]int64, start, end int) {
	//确认终止条件，否则将无限递归下去
	if start >= end {
		return
	}
	pivotIdx := (start + end) / 2
	pivotV := (*arr)[pivotIdx]
	l, r := start, end
	for l < r {
		for (*arr)[l] < pivotV {
			l++
		}
		for (*arr)[r] > pivotV {
			r--
		}
		//缺少了这三行导致失败的数组:45 12 33 10 44 0 27 27 20
		if l >= r { //课堂上我们没有在这里break导致意外的排序失败
			break //课堂上我们没有在这里break导致意外的排序失败
		} //课堂上我们没有在这里break导致意外的排序失败
		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
	}
	fmt.Println("l:", l, "r", r)
	fmt.Println(*arr)
	if l == r {
		l++
		r--
	}
	if r > start {
		quickSort(arr, start, r)
	}
	if l < end {
		quickSort(arr, l, end)
	}
}

func main() {
	arrSize := 10000000
	arr := []int64{}
	//for i := 0; i < arrSize; i++ {
	//	arr = append(arr, rand.Intn(50))
	//}
	for i := 0; i < arrSize; i++ {
		all, _ := rand.Int(rand.Reader, big.NewInt(3000))
		arr = append(arr, all.Int64())
	}
	fmt.Println(arr)
	quickSort(&arr, 0, arrSize-1)
	fmt.Println(arr)
}
