package main

import "fmt"

func main() {
	var money = 400
	var busy bool = true
	switch {
	case money <= 20:
		fmt.Println("点个外卖")
		if busy == false {
			fmt.Println("再吃个零食")
		}
	case money <= 200 && money > 20:
		fmt.Println("下个馆子")
		if busy {
			fmt.Println("再吃个水果")
		}
	default:
		fmt.Println("再想想")
	}
	fmt.Println("end")
}
