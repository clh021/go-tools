package grpcWebClient

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "test/service/grpcWeb/time/grpcWebClient/time/v1"
)

const (
	listenAddress = "0.0.0.0:9090"
)

type timeService struct {
}

func (t *timeService) GetCurrentTime(ctx context.Context, req *pb.GetCurrentTimeRequest) (*pb.GetCurrentTimeResponse, error) {
	log.Println("Got time request")
	return &pb.GetCurrentTimeResponse{CurrentTime: time.Now().String()}, nil
}

func Main() {
	log.Printf("Time service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(s, &timeService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
