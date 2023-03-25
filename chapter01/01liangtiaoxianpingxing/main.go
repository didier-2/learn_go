package main

import "fmt"

func main() {
	var p1x, p1y, p2x, p2y, p3x, p3y, p4x, p4y float64
	p1x, p1y = getPointXYFromInput()
	p2x, p2y = getPointXYFromInput()
	p3x, p3y = getPointXYFromInput()
	p4x, p4y = getPointXYFromInput()

	//fmt.Print("录入第一个点x")
	//fmt.Scanln(&p1x)
	//fmt.Print("录入第一个点y")
	//fmt.Scanln(&p1y)
	//fmt.Print("录入第二个点x")
	//fmt.Scanln(&p2x)
	//fmt.Print("录入第二个点y")
	//fmt.Scanln(&p2y)
	//fmt.Print("录入第三个点x")
	//fmt.Scanln(&p3x)
	//fmt.Print("录入第三个点y")
	//fmt.Scanln(&p3y)
	//fmt.Print("录入第四个点x")
	//fmt.Scanln(&p4x)
	//fmt.Print("录入第四个点y")
	//fmt.Scanln(&p4y)

	k1 := calculateK(p2y, p1y, p4x, p3x)
	k2 := calculateK(p4y, p3y, p2x, p1x)

	kSame(k1, k2)
}

func getPointXYFromInput() (float64, float64) {
	var x, y float64
	fmt.Print("录入x")
	fmt.Scanln(&x)
	fmt.Print("录入y")
	fmt.Scanln(&y)
	return x, y
}
func calculateK(y2, y1, x4, x3 float64) float64 {

	return (y2 - y1) * (x4 - x3)
}
func kSame(k1, k2 float64) bool {
	var s bool //先定义一个变量接收true、false
	if k1 == k2 {
		s = true
		fmt.Println("两条线平行")
	} else {
		s = false
		fmt.Println("两条线不平行")
	}
	return s
}
