package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	var input Interface = &fakeInterface{
		name:       "小强",
		sex:        "男",
		baseWeight: 71,
		baseTall:   1.7,
		baseAge:    35,
	}

	for {
		func() {
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				log.Fatal("拨号失败：", err)
			}
			defer conn.Close()
			fmt.Println("连接成功，开始聊天")
			pi, err := input.ReadPersonalInformation()
			if err != nil {
				log.Printf("Warning:读取失败，请重新录入：", err)
				return
			}
			data, err := json.Marshal(pi)
			if err != nil {
				log.Printf("无法编码个人信息", err)
				return
			}

			log.Println("读取到数据：", string(data))
			talk(conn, string(data))

		}()
		time.Sleep(1 * time.Second)
	}
	//talk(conn, "你好")
	//talk(conn, "你好")
	//talk(conn, "你好")
	//talk(conn, "你是谁")
	//talk(conn, "再见")

}

func talk(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Printf("发送消息失败：", err)
	} else {
		data := make([]byte, 1024)
		validLen, err := conn.Read(data)
		if err != nil {
			log.Printf("warning:读取服务器返回数据时出错:", err)
		} else {
			validData := data[:validLen]
			log.Println("去", message, "回", string(validData))
		}
	}
}
