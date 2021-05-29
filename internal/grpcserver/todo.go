package grpcserver

import (
	"context"
	"fmt"
	"github.com/peppys/service-template/gen/go/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type repository interface {
	FindAll(ctx context.Context) ([]*proto.Todo, error)
	Create(ctx context.Context, text string, author string) (*proto.Todo, error)
	FindById(ctx context.Context, author string) (*proto.Todo, error)
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

	return &proto.ListAllResponse{
		Todos: records,
	}, nil
}

func (s *TodoGrpcServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.Todo, error) {
	record, err := s.repository.Create(ctx, request.GetText(), request.GetAuthor())
	if err != nil {
		return nil, fmt.Errorf("error while creating: %v", err)
	}

	return record, nil
}

func (s *TodoGrpcServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.Todo, error) {
	record, err := s.repository.FindById(ctx, request.GetId())
	if err != nil {
		return nil, fmt.Errorf("error while finding: %v", err)
	}

	return record, nil
}
