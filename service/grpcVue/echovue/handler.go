package echovue

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
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
		log.Printf(" streaming resp %d", i)
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
		log.Printf(" streaming resp %d", i)
		resp := &ServerStreamingEchoResponse{Message: fmt.Sprintf("ServerStreamingEchoAbort resp %d", i)}
		if err := streamAbort.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) ClientStreamingEcho(stream EchoService_ClientStreamingEchoServer) error {
	log.Printf("Received new ClientStreamingEcho request")
	return nil
}

func (s *Server) FullDuplexEcho(fullDuplex EchoService_FullDuplexEchoServer) error {
	log.Printf("Received new FullDuplexEcho request")
	return nil
}
func (s *Server) HalfDuplexEcho(halfDuplex EchoService_HalfDuplexEchoServer) error {
	log.Printf("Received new HalfDuplexEcho request")
	return nil
}

// func (s *Server) GetTodos(ctx context.Context, _ *GetTodoParams) (*TodoResponse, error) {
// 	log.Printf("get tasks")
// 	return &TodoResponse{Todos: s.Todos}, nil
// }

// func (s *Server) DeleteTodo(ctx context.Context, delTodo *DeleteTodoParams) (*DeleteResponse, error) {
// 	var updatedTodos []*TodoObject
// 	for index, todo := range s.Todos {
// 		if todo.Id == delTodo.Id {
// 			updatedTodos = append(s.Todos[:index], s.Todos[index+1:]...)
// 			break
// 		}
// 	}
// 	s.Todos = updatedTodos
// 	return &DeleteResponse{Message: "success"}, nil
// }
