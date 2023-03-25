package main

import "fmt"

func main() {
	var money = 19
	var busy bool
	switch {
	case money <= 20 && money >= 0:
		fmt.Println("点个外卖")
		if busy {
			break
		}
		fmt.Println("再吃个水果")
		fallthrough
	case money > 20 && money <= 200:
		fmt.Println("下个馆子")
		if busy {
			break
		}
		fmt.Println("吃个零食")
	default:
		fmt.Println("再想想")
	}

}
