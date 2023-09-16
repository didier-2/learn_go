package main

import (
	"context"
	"fmt"
	"go.learn/pkg/apis"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	c := apis.NewRankServiceClient(conn)
	w, err := c.WatchPersons(context.TODO(), &apis.Null{})
	if err != nil {
		log.Fatal("启动watcher时失败:", err)
	}
	for {
		pi, err := w.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("服务器告知说完了")
				break
			}
			log.Fatal("接受异常:", err)
		}
		log.Println("收到新变动：", pi.String())
	}
	ret, err := c.Register(context.TODO(), &apis.PersonalInformation{Name: "tom"})
	if err != nil {
		log.Fatal("注册失败", err)
	}
	log.Println("注册成功", ret)
	log.Println("开始批量注册")
	regCli, err := c.RegisterPersons(context.TODO())
	if err != nil {
		log.Fatal("获取批量注册客户端失败", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败", err)
	}
	resp, err := regCli.CloseAndRecv()
	if err != nil {
		log.Fatal("无法接受结果", err)
	}
	log.Println("批量注册结果：", resp.String())
}
