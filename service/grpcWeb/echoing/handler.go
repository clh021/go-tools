package echoing

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	log.Printf("Received new echo request %s", req)
	echoObject := &EchoResponse{
		Message:      "receive:Echo:",
		MessageCount: 1,
	}
	return echoObject, nil
}

func (s *Server) EchoAbort(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return &EchoResponse{
		Message:      "receive:EchoAbort:",
		MessageCount: 2,
	}, nil
}
func (s *Server) NoOp(ctx context.Context, req *Empty) (*Empty, error) {
	return &Empty{}, nil
}
func (s *Server) ServerStreamingEcho(ctx *ServerStreamingEchoRequest, stream EchoService_ServerStreamingEchoServer) error {
	return nil
}
func (s *Server) ServerStreamingEchoAbort(ctx *ServerStreamingEchoRequest, streamAbort EchoService_ServerStreamingEchoAbortServer) error {
	return nil
}
func (s *Server) ClientStreamingEcho(stream EchoService_ClientStreamingEchoServer) error {
	return nil
}

func (s *Server) FullDuplexEcho(fullDuplex EchoService_FullDuplexEchoServer) error {
	return nil
}
func (s *Server) HalfDuplexEcho(halfDuplex EchoService_HalfDuplexEchoServer) error {
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
