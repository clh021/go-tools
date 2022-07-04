package echovue

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	log.Printf("Received new echo request %s", req)
	echoObject := &EchoResponse{
		Message:      fmt.Sprintf("receive:Echo: %s", req),
		MessageCount: 1,
	}
	return echoObject, nil
}

func (s *Server) EchoAbort(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	log.Printf("Received new EchoAbort request %s", req)
	return &EchoResponse{
		Message:      fmt.Sprintf("receive:EchoAbort: %s", req),
		MessageCount: 2,
	}, nil
}

func (s *Server) NoOp(ctx context.Context, req *Empty) (*Empty, error) {
	log.Printf("Received new NoOp request %s", req)
	return &Empty{}, nil
}

func (s *Server) ServerStreamingEcho(ctx *ServerStreamingEchoRequest, stream EchoService_ServerStreamingEchoServer) error {
	log.Printf("Received new ServerStreamingEcho request")
	for i := 0; i < 10; i++ {
		log.Printf(" ServerStreamingEcho resp %d", i)
		resp := &ServerStreamingEchoResponse{Message: fmt.Sprintf("ServerStreamingEcho resp %d", i)}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) ServerStreamingEchoAbort(ctx *ServerStreamingEchoRequest, streamAbort EchoService_ServerStreamingEchoAbortServer) error {
	log.Printf("Received new ServerStreamingEchoAbort request")
	for i := 0; i < 10; i++ {
		log.Printf(" ServerStreamingEchoAbort resp %d", i)
		resp := &ServerStreamingEchoResponse{Message: fmt.Sprintf("ServerStreamingEchoAbort resp %d", i)}
		if err := streamAbort.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) ClientStreamingEcho(stream EchoService_ClientStreamingEchoServer) error {
	log.Printf("Received new ClientStreamingEcho request")
	// clientStreamReq, err := stream.Recv()
	// if err != nil {
	// 	return err
	// }
	// log.Printf(" ClientStreamingEcho resp %s", clientStreamReq.GetMessage())

	ctx := stream.Context()
	var msgCount int32 = 0
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			msgCount++
			// 接收从客户端发来的消息
			clientStreamReq, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				return err
			}
			log.Printf(" ClientStreamingEcho resp1 %s", clientStreamReq.GetMessage())
			log.Printf(" ClientStreamingEcho resp2 %s", clientStreamReq.String())

			resp := &ClientStreamingEchoResponse{MessageCount: msgCount}
			if err := stream.SendAndClose(resp); err != nil {
				return err
			}
		}
	}
}

func (s *Server) FullDuplexEcho(fullDuplex EchoService_FullDuplexEchoServer) error {
	log.Printf("Received new FullDuplexEcho request")
	return nil
}
func (s *Server) HalfDuplexEcho(halfDuplex EchoService_HalfDuplexEchoServer) error {
	log.Printf("Received new HalfDuplexEcho request")
	return nil
}
