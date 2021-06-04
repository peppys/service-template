package grpcservers

import (
	"context"
	"github.com/google/uuid"
	"github.com/peppys/service-template/gen/go/proto"
	"github.com/peppys/service-template/internal/entities"
	"github.com/peppys/service-template/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type repository interface {
	FindAllWhere(context.Context, entities.Todo) ([]*entities.Todo, error)
	FindFirstWhere(context.Context, entities.Todo) (*entities.Todo, error)
	Create(context.Context, *entities.Todo) (*entities.Todo, error)
	DeleteByID(context.Context, string) error
}
type TodoGrpcServer struct {
	proto.UnimplementedTodoServiceServer
	repository
}

func NewTodoGrpcServer(repo repository) *TodoGrpcServer {
	return &TodoGrpcServer{repository: repo}
}

func (s *TodoGrpcServer) ListAll(ctx context.Context, request *emptypb.Empty) (*proto.ListAllResponse, error) {
	u, err := utils.UserClaimsFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	records, err := s.repository.FindAllWhere(ctx, entities.Todo{
		UserID: u.UUID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
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
	u, err := utils.UserClaimsFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	record, err := s.repository.Create(ctx, &entities.Todo{
		Text:   request.GetText(),
		UserID: u.UUID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return toProto(record), nil
}

func (s *TodoGrpcServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.Todo, error) {
	u, err := utils.UserClaimsFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	recordUuid, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	record, err := s.repository.FindFirstWhere(ctx, entities.Todo{
		ID:     recordUuid,
		UserID: u.UUID,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "todo not found")
	}

	return toProto(record), nil
}

func (s *TodoGrpcServer) Delete(ctx context.Context, request *proto.DeleteRequest) (*emptypb.Empty, error) {
	u, err := utils.UserClaimsFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	recordUuid, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	_, err = s.repository.FindFirstWhere(ctx, entities.Todo{
		ID:     recordUuid,
		UserID: u.UUID,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "todo not found")
	}
	err = s.repository.DeleteByID(ctx, recordUuid.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func toProto(todo *entities.Todo) *proto.Todo {
	return &proto.Todo{
		Id:        todo.ID.String(),
		Text:      todo.Text,
		UserId:    todo.UserID.String(),
		Timestamp: todo.CreatedAt.String(),
	}
}
