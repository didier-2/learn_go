package main

import "fmt"

func main() {
	a := "您好"
	fmt.Println(a)
	fmt.Println("byte a:", []byte(a), len([]byte(a)))
	fmt.Println("rune a:", []rune(a), len([]rune(a)))

	aBytes := []rune(a)
	//aBytes := []byte(a)//[]int 不能转，[]byte 可以和string互通
	fmt.Println(aBytes)
	fmt.Println("修改切片的内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a)
}
