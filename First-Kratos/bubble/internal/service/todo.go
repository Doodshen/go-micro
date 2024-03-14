package service

import (
	"context"
	"errors"

	pb "bubble/api/bubble/v1"
	"bubble/internal/biz"
)

// 定义服务模型
type TodoService struct {
	pb.UnimplementedTodoServer

	uc *biz.TodoUsecase
}

// mustEmbedUnimplementedGreeterServer implements v1.GreeterServer.
func (s *TodoService) mustEmbedUnimplementedGreeterServer() {
	panic("unimplemented")
}

// SayHello implements v1.GreeterServer.
func (s *TodoService) SayHello(context.Context, *v1.HelloRequest) (*v1.HelloReply, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedGreeterServer implements v1.GreeterServer.
func (s *TodoService) mustEmbedUnimplementedGreeterServer() {
	panic("unimplemented")
}

func NewTodoService(uc *biz.TodoUsecase) *TodoService {
	return &TodoService{
		uc: uc,
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	//请求来了
	if len(req.GetTitle()) == 0 {
		return &pb.CreateTodoReply{}, errors.New("invalid params ")
	}
	//调用业务逻辑
	data, err := s.uc.Create(ctx, &biz.Todo{Title: req.Title})
	if err != nil {
		return &pb.CreateTodoReply{
			Id:     data.ID,
			Title:  data.Title,
			Status: data.Status,
		}, nil
	}
	//返回响应
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
