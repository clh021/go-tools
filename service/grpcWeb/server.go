package grpcWeb

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"runtime/debug"
	"test/service/grpcWeb/echoing"
)

func Main() {

	// 退出自动恢复服务
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s\n\n%s\n", r, debug.Stack())
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 18080))
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
