package main

import (
	"context"
	"go.learn/pkg/apis"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	startGRPCServer(ctx)
}

func startGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)
	apis.RegisterRankServiceServer(s, &rankServer{
		persons: map[string]*apis.PersonalInformation{},
	})
	go func() {
		select {
		case <-ctx.Done():
			s.Stop()
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server:%v", err)
	}
}
