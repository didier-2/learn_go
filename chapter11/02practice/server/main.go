package main

import (
	"encoding/json"
	"flag"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"go.learn/pkg/apis"
	"log"
	"net"
	"time"
)

//
//var qa = map[string]string{
//	"你好":     "你好",
//	"你是谁":    "我是小s",
//	"你是男是女":  "你猜",
//	"今天天气如何": "今天天气不错",
//	"再见":     "再见",
//}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	rank := NewFatRateRank()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("warning:建立连接失败：", err)
			continue
		}
		fmt.Println(conn)
		go talk(conn, rank)
	}
}

func talk(conn net.Conn, rank *FatRateRank) {
	defer fmt.Printf("结束链接", conn)
	defer conn.Close()
	for {
		finalReceivedMessage := make([]byte, 0, 1024)

		encounterError := false
		for {
			buf := make([]byte, 1024)
			valid, err := conn.Read(buf)
			if err != nil {
				log.Println("warning:读取数据时出问题：", err)
				log.Println("重新读取", err)
				encounterError = true
				time.Sleep(1 * time.Second)
				break
			}
			if valid == 0 {
				break
			}

			validContent := buf[:valid]
			finalReceivedMessage = append(finalReceivedMessage, validContent...)
			if valid < len(buf) {
				break
			}
		}
		if encounterError {
			conn.Write([]byte(`服务器获取数据失败，重新输入`)) //handle error
			continue
		}

		pi := &apis.PersonalInformation{}
		if err := json.Unmarshal(finalReceivedMessage, pi); err != nil {
			conn.Write([]byte(`服务器获取数据失败，重新输入`)) //handle error
			continue
		}

		bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
		if err != nil {
			conn.Write([]byte(`无法计算你的bmi，重新输入`)) //handle error
			continue
		}
		fr := gobmi.CalcFatRate(bmi, int(pi.Age), pi.Sex)

		rank.inputRecord(pi.Name, fr)
		rankId, _ := rank.getRand(pi.Name)

		conn.Write([]byte(fmt.Sprintf("姓名:%s,BMI:%v,体脂率:%v,排名:%v", pi.Name, bmi, fr, rankId))) //handle error
		break

	}
}
