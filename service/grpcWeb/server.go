package grpcWeb

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"test/service/grpcWeb/echoing"
)

func Main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := echoing.Server{}
	grpcServer := grpc.NewServer()

	echoing.RegisterEchoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}
