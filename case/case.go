package main

import "fmt"

func main() {
	var money = 20
	switch money {
	case 20:
		fmt.Println("点外卖")
	case 200:
		fmt.Println("下馆子")
	case 2000:
		fmt.Println("去米其林")
	default:
		fmt.Println("再想想")
	}
	fmt.Println("end")
}
