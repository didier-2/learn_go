package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := "D:/Desktop/key.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("无法打开文件", filepath, "错误信息是", err)
		os.Exit(1)
	}
	defer file.Close()

	b := make([]byte, 10)
	var n int
	for i := 0; i < 2; i++ {
		n, err = file.Read(b)
		if err != nil {
			fmt.Println("无法读取文件", err)
			os.Exit(2)
		}
		//fmt.Println("读出的内容是：", string(b)) //如果不转换数据类型，那么读出的是一长串数据
		fmt.Println("读出的内容是：", string(b[:n])) //记得给后续数据程序使用时，截取实际得到的数据，而不是全部。否则后续的程序会把无效的读取也作为正常的数据处理
		fmt.Println("n的大小是：", n)
		file.Seek(0, io.SeekStart)   //SeekStart表示从开始数数，转移游标
		file.Seek(3, io.SeekCurrent) //SeekCurrent表示从目标游标位置开始，转移到相对位置
		//todo handle error
	}

}
