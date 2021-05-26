package grpcserver

import (
	"context"
	"fmt"
	"github.com/peppys/service-template/gen/go/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type repository interface {
	FindAll(ctx context.Context) ([]*todo.Todo, error)
	Create(ctx context.Context, text string, author string) (*todo.Todo, error)
	FindById(ctx context.Context, author string) (*todo.Todo, error)
}
type TodoGrpcServer struct {
	todo.UnimplementedTodoServiceServer
	repository
}

func NewTodoGrpcServer(repo repository) *TodoGrpcServer {
	return &TodoGrpcServer{repository: repo}
}

func (s *TodoGrpcServer) ListAll(ctx context.Context, request *emptypb.Empty) (*todo.ListAllResponse, error) {
	records, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while listing: %v", err)
	}

	return &todo.ListAllResponse{
		Todos: records,
	}, nil
}

func (s *TodoGrpcServer) Create(ctx context.Context, request *todo.CreateRequest) (*todo.Todo, error) {
	record, err := s.repository.Create(ctx, request.GetText(), request.GetAuthor())
	if err != nil {
		return nil, fmt.Errorf("error while creating: %v", err)
	}

	return record, nil
}

func (s *TodoGrpcServer) Get(ctx context.Context, request *todo.GetRequest) (*todo.Todo, error) {
	record, err := s.repository.FindById(ctx, request.GetId())
	if err != nil {
		return nil, fmt.Errorf("error while finding: %v", err)
	}

	return record, nil
}
