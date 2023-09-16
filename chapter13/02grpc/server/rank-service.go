package main

import (
	"context"
	"go.learn/chapter13/02grpc/server/frinterface"
	"go.learn/pkg/apis"
	"io"
	"log"
	"sync"
)

var _ apis.RankServiceServer = &rankServer{}

type rankServer struct {
	sync.Mutex
	//persons map[string]*apis.PersonalInformation
	rankS    frinterface.ServerInterface
	personCh chan *apis.PersonalInformation
}

func (r rankServer) Update(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	r.regPerson(information)
	return r.rankS.UpdatePersonalInformation(information)
}

func (r rankServer) GetFatRate(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalRank, error) {
	return r.rankS.GetFatRate(information.Name)
}

func (r rankServer) GetTop(ctx context.Context, null *apis.Null) (*apis.Ranks, error) {
	top, err := r.rankS.GetTop()
	if err != nil {
		log.Println("获取榜单时出错：", err)
		return nil, err
	}
	return &apis.Ranks{PersonalRanks: top}, nil
}

func (r rankServer) regPerson(pi *apis.PersonalInformation) {
	r.rankS.RegisterPersonalInformation(pi) //todo handle error
	r.personCh <- pi
}

func (r rankServer) WatchPersons(null *apis.Null, server apis.RankService_WatchPersonsServer) error {
	for pi := range r.personCh {
		if err := server.Send(pi); err != nil {
			log.Println("发送失败，结束:", err)
			return err
		}
	}
	return nil
}

func (r rankServer) RegisterPersons(server apis.RankService_RegisterPersonsServer) error {
	pis := &apis.PersonalInformationList{}
	for {
		pi, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Printf("Warning:获取人员注册时失败：%v\n", err)
			return err
		}
		pis.Items = append(pis.Items, pi)
		r.regPerson(pi)
	}
	log.Println("连续得到注册清单:", pis.String())
	return server.SendAndClose(pis)
}

func (r rankServer) Register(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.regPerson(information)
	log.Printf("收到新注册人：%s\n", information.String())
	return information, nil
}
