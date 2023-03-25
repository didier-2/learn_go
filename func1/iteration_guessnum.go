package main

import "fmt"

func guess(left, right uint) {
	guessed := (left + right) / 2
	var getFromInput string
	fmt.Println("我猜的是", guessed)
	fmt.Print("如果高了 ，输入1；如果低了，输入0；对了，输入9")
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case "1":
		if left == right {
			fmt.Println("你是不是改变主意了")
			return
		}
		guess(left, guessed-1)
	case "2":
		if left == right {
			fmt.Println("你是不是改变主意了")
			return
		}
		guess(guessed+1, right)
	case "9":
		fmt.Println("猜对了，你心里想的数字是：", guessed)
	}
}
