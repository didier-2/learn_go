package main

import "fmt"

func main() {
	xdInfo := [3]string{"小丁", "小黄", "小彬"}
	fmt.Println(xdInfo)
	for i := 0; i < len(xdInfo)/2; i++ { //数组反转
		temp := xdInfo[i]
		xdInfo[i] = xdInfo[len(xdInfo)-i-1]
		xdInfo[len(xdInfo)-i-1] = temp
	}
	fmt.Println(xdInfo)

	fmt.Println("*********************")
	newPersonInfos := [3][3]string{
		[3]string{"小丁", "小黄", "小彬"},
		[3]string{"小丁", "小黄", "小彬"},
		[3]string{"小丁", "小黄", "小彬"},
	}
	for _, val := range newPersonInfos {
		fmt.Println(val)
	}

	//...支持动态添加
	newPersonInfos2 := [...][3]string{
		[3]string{"小丁", "小黄", "小彬"},
		[3]string{"小丁", "小黄", "小彬"},
		[3]string{"小丁", "小黄", "小彬"},
		[3]string{"小丁", "小黄", "小彬"},
		[3]string{"小丁", "小黄", "小彬"},
	}
	for _, val := range newPersonInfos2 {
		fmt.Println(val)
	}

}
