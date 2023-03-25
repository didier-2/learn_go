package main

import "fmt"

func panicAndRecover() {
	defer coverPanicUpgraded() //这样能抓住严重错误
	defer coverPanic()
	var nameScore map[string]int = nil
	nameScore["小丁"] = 100
}

func coverPanic() { //未能抓住panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("程序发生严重错误", r)
		}
	}()
}
func coverPanicUpgraded() {
	if r := recover(); r != nil {
		fmt.Println("程序发生严重错误", r)
	}
}
