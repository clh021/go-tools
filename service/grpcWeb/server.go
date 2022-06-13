package grpcWeb

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"test/service/grpcWeb/todo"
)

func Main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50096))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := todo.Server{}
	grpcServer := grpc.NewServer()

	todo.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}
