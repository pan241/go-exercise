package service

import (
	"context"

	pb "bubble/api/bubble/v1"
)

type TodoService struct {
	pb.UnimplementedTodoServer
}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {

	return &pb.CreateTodoReply{}, nil
}
func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoReply, error) {
	return &pb.UpdateTodoReply{}, nil
}
func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	return &pb.DeleteTodoReply{}, nil
}
func (s *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoReply, error) {
	return &pb.GetTodoReply{}, nil
}
func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	return &pb.ListTodoReply{}, nil
}
