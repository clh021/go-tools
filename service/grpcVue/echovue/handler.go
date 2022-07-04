package echovue

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

// 一个请求，一个响应
// 服务器按原样返回客户端消息。
func (s *Server) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	log.Printf("Received new echo request %s", req)
	echoObject := &EchoResponse{
		Message:      fmt.Sprintf("receive:Echo: %s", req),
		MessageCount: 1,
	}
	return echoObject, nil
}

// 发回中止状态。
func (s *Server) EchoAbort(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	log.Printf("Received new EchoAbort request %s", req)
	return &EchoResponse{
		Message:      fmt.Sprintf("receive:EchoAbort: %s", req),
		MessageCount: 2,
	}, nil
}

// 一个空请求，零处理，后接一个空响应
//（做消息序列化的最小努力）。
func (s *Server) NoOp(ctx context.Context, req *Empty) (*Empty, error) {
	log.Printf("Received new NoOp request %s", req)
	return &Empty{}, nil
}

// 一个请求后跟一系列响应（流式下载）。
// 服务器会重复返回相同的客户端消息。
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

// 一个请求后跟一系列响应（流式下载）。
// 服务器直接中止。
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

// 一系列请求，然后是一个响应（流式上传）。
// 服务器返回消息的总数作为结果。
func (s *Server) ClientStreamingEcho(stream EchoService_ClientStreamingEchoServer) error {
	log.Printf("Received new ClientStreamingEcho request")
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

// 服务器立即回显每条消息的一系列请求。
// 服务器按顺序返回相同的客户端消息。
// 例如，这就是语音 API 的工作方式。
func (s *Server) FullDuplexEcho(fullDuplex EchoService_FullDuplexEchoServer) error {
	log.Printf("Received new FullDuplexEcho request")
	return nil
}

// 一个请求序列，后面跟着一个响应序列。
// 服务器缓存所有客户端消息，然后返回相同的
// 客户端半关闭流后，客户端消息一一对应。
// 这就是图像识别 API 的工作方式。
func (s *Server) HalfDuplexEcho(halfDuplex EchoService_HalfDuplexEchoServer) error {
	log.Printf("Received new HalfDuplexEcho request")
	return nil
}
