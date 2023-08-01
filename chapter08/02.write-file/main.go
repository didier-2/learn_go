package main

import (
	"fmt"
	"os"
)

func main() {
	filepath := "D:/Desktop/小丁.txt"
	//writeFile(filepath)
	writeFilewithAppend(filepath)

}

func writeFile(filepath string) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("无法打开文件", filepath, "错误信息是", err)
		os.Exit(1)
	}
	defer file.Close()

	//b := make([]byte, 10)

	//var n int
	_, err = file.Write([]byte("this is first line\n"))
	fmt.Println(err)
	_, err = file.Write([]byte("this is first line\n"))
	fmt.Println(err)
	_, err = file.WriteString("第二行内容\n")
	fmt.Println(err)
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	fmt.Println(err)
}
func writeFilewithAppend(filepath string) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //linux 权限
	if err != nil {
		fmt.Println("无法打开文件", filepath, "错误信息是", err)
		os.Exit(1)
	}
	defer file.Close()

	//b := make([]byte, 10)

	//var n int
	_, err = file.Write([]byte("this is first line\n"))
	fmt.Println(err)
	_, err = file.Write([]byte("this is first line\n"))
	fmt.Println(err)
	_, err = file.WriteString("第二行内容\n")
	fmt.Println(err)
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	fmt.Println(err)
}
