package main

import "fmt"

func main() {
	msg := constructHalloMessage("小丁")
	fmt.Println(msg)

}
func constructHalloMessage(name string) string {
	//fmt.Println("你好", name)
	return "你好" + name
}
func caleBMI(tall float64, weight float64) float64 {
	return tall / (weight * weight)
}
