package main

import "fmt"

func main() {
	var money interface{} = 12.0
	switch newmoney := money.(type) {
	case int:
		fmt.Println("int", newmoney+11)
	case int64:
		fmt.Println("int64", newmoney+12)
	case string:
		fmt.Println("String", newmoney+"s")
	case float32:
		fmt.Println("float32", newmoney+13)
	case float64:
		fmt.Println("float64", newmoney+14)
	default:
		fmt.Println("是未知类型")
	}
	fmt.Println("end", money)
}
