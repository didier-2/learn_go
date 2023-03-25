package main

import "fmt"

func main() {
	var m1 map[string]int = nil
	//m1["a"]=1
	fmt.Println("m1没有实例化，直接取数：", m1)
	m2 := map[string]int{}
	m3 := map[string]int{"小丁": 92, "小黄": 90, "小彬": 100}
	fmt.Println(m1, m2, m3)

	xdScore, ok := m3["小韦"]
	fmt.Println(xdScore, ">>>>>>>", ok)

	//添加
	m3["小韦"] = 95
	fmt.Println(m3)
	xdScore, ok = m3["小韦"]
	fmt.Println(xdScore, ">>>>>>>", ok)
	//删除
	delete(m3, "小韦")
	fmt.Println(m3)
	//修改
	m3["小彬"] = 95
	fmt.Println(m3)
	//遍历
	for name, score := range m3 {
		fmt.Printf("%s\t%d\n", name, score)
	}

}
