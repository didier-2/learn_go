//package main
//
//import "fmt"
//
//func main() {
//	var money = 20
//	if money <= 20 {
//		fmt.Println("点外卖")
//	} else if money <= 200 {
//		fmt.Println("下馆子")
//	} else if money <= 2000 {
//		fmt.Println("吃米其林")
//	} else {
//		fmt.Println("再想想")
//	}
//
//}

package main

import "fmt"

func main() {
	var money = 20
	if money <= 20 {
		fmt.Println("点外卖")
	}
	if money <= 200 {
		fmt.Println("下个馆子")
	}
	if money < 2000 {
		fmt.Println("吃大餐")
	}

}
