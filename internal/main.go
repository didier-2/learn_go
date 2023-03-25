package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(finishTime, startTime)
		fmt.Println(finishTime.Sub(startTime))
	}()
	arr := make([]string, 4, 4)
	fmt.Println("len:", len(arr), "cap:", cap(arr))
	fmt.Println("default:", arr[0])
	fmt.Println("default:", arr[1])
	fmt.Println("default:", arr[2])

	arr2 := make([]string, 3, 4)
	copy(arr, arr2)

	i := 333
	j := &i
	fmt.Println(reflect.TypeOf(j))

}
