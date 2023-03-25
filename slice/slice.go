package main

import "fmt"

func main() {
	a := []int{}
	c := []int{111, 222, 333, 444, 555}
	fmt.Println("a是切片，能追加元素")
	a = append(a, 333) //追加元素
	fmt.Println(a)
	b := [0]int{}
	fmt.Println(b)

	xdInfo := []string{"小丁", "男", "23"}
	fmt.Println(xdInfo)
	for i, v := range xdInfo {
		fmt.Println(i, v)
	}
	fmt.Println(xdInfo[0])
	xdInfo = append(xdInfo[:1], xdInfo[2:]...) //切片的删除
	fmt.Println(xdInfo)

	fmt.Println("切片的插入")
	//切片的插入操作
	backup := append([]int{}, c[1:]...)
	c = append(c[:1], 999)
	c = append(c, backup...)
	fmt.Println(c)

}

//package main
//
//import "fmt"
//
//func main() {
//	a := [3][3]int{}
//	a[0] = [3]int{11, 12, 13}
//	a[1] = [3]int{21, 22, 23}
//	a[2] = [3]int{31, 32, 33}
//	fmt.Println(a)
//	fmt.Println("遍历一层")
//	for i := 0; i < 3; i++ {
//		fmt.Println(a[i])
//	}
//	fmt.Println("遍历两层")
//	for i, val := range a {
//		for j := 0; j < len(val); j++ {
//			fmt.Printf("[%d:%d]%d\t", i, j, val[j])
//		}
//		fmt.Println()
//	}
//}
