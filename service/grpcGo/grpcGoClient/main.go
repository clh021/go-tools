package grpcGoClient

import (
	"context"
	"flag"
	"log"
	"time"

	pb "test/service/grpcGo/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

func Main() {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	name := flag.String("name", defaultName, "Name to greet")
	flag.Parse()

	// 初始化 grpc 连接
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 建立上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用 grpc 接口
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
