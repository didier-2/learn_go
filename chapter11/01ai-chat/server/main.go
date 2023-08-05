package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var qa = map[string]string{
	"你好":     "你好",
	"你是谁":    "我是小s",
	"你是男是女":  "你猜",
	"今天天气如何": "今天天气不错",
	"再见":     "再见",
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("warning:建立连接失败：", err)
			continue
		}
		fmt.Println(conn)
		go talk(conn)
	}
}

func talk(conn net.Conn) {
	defer fmt.Printf("结束链接", conn)
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("warning:建立连接失败：", err)
			continue
		}
		coutent := buf[:valid]
		resp, ok := qa[string(coutent)]
		if !ok {
			log.Println("没找到回答，问他说了啥")
			conn.Write([]byte(`我听不懂你在说什么`)) //handle error
			continue
		}
		conn.Write([]byte(resp)) //handle error

		if string(coutent) == "再见" {
			break
		}
	}
}
