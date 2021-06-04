package grpcservers

import (
	"context"
	"fmt"
	"github.com/peppys/service-template/gen/go/proto"
	"github.com/peppys/service-template/internal/entities"
	"google.golang.org/protobuf/types/known/emptypb"
)

type repository interface {
	FindAll(ctx context.Context) ([]*entities.Todo, error)
	Create(ctx context.Context, todo *entities.Todo) (*entities.Todo, error)
	FindById(ctx context.Context, id string) (*entities.Todo, error)
}
type TodoGrpcServer struct {
	proto.UnimplementedTodoServiceServer
	repository
}

func NewTodoGrpcServer(repo repository) *TodoGrpcServer {
	return &TodoGrpcServer{repository: repo}
}

func (s *TodoGrpcServer) ListAll(ctx context.Context, request *emptypb.Empty) (*proto.ListAllResponse, error) {
	records, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while listing: %v", err)
	}

	responses := make([]*proto.Todo, 0, len(records))
	for _, record := range records {
		responses = append(responses, toProto(record))
	}
	return &proto.ListAllResponse{
		Todos: responses,
	}, nil
}

func (s *TodoGrpcServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.Todo, error) {
	record, err := s.repository.Create(ctx, &entities.Todo{
		Text:   request.GetText(),
		Author: request.GetAuthor(),
	})
	if err != nil {
		return nil, fmt.Errorf("error while creating: %v", err)
	}

	return toProto(record), nil
}

func (s *TodoGrpcServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.Todo, error) {
	record, err := s.repository.FindById(ctx, request.GetId())
	if err != nil {
		return nil, fmt.Errorf("error while finding: %v", err)
	}

	return toProto(record), nil
}

func toProto(todo *entities.Todo) *proto.Todo {
	return &proto.Todo{
		Id:        todo.ID.String(),
		Text:      todo.Text,
		Author:    todo.Author,
		Timestamp: todo.CreatedAt.String(),
	}
}
