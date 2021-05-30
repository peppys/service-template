package grpcserver

import (
	"context"
	"fmt"
	"github.com/peppys/service-template/gen/go/proto"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthGrpcServer struct {
	grpc_health_v1.UnimplementedHealthServer
	proto.UnimplementedHealthServiceServer
}

func NewHealthGrpcServer() *HealthGrpcServer {
	return &HealthGrpcServer{}
}

func (s *HealthGrpcServer) Liveness(ctx context.Context, empty *emptypb.Empty) (*proto.LivenessResponse, error) {
	return &proto.LivenessResponse{
		Ok: true,
	}, nil
}

func (s *HealthGrpcServer) Readiness(ctx context.Context, empty *emptypb.Empty) (*proto.ReadinessResponse, error) {
	return &proto.ReadinessResponse{
		Ok: true,
		Ready: &proto.ReadinessResponse_DependencyReadiness{
			Datastore: true,
		},
	}, nil
}

func (s *HealthGrpcServer) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	status, err := s.checkReadinessAndMapToHealthCheckStatus(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while checking health: %v", err)
	}

	return &grpc_health_v1.HealthCheckResponse{
		Status: status,
	}, nil
}

func (s *HealthGrpcServer) Watch(request *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	status, err := s.checkReadinessAndMapToHealthCheckStatus(context.Background())
	if err != nil {
		return fmt.Errorf("error while checking health: %v", err)
	}

	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: status,
	})
}

func (s *HealthGrpcServer) checkReadinessAndMapToHealthCheckStatus(ctx context.Context) (grpc_health_v1.HealthCheckResponse_ServingStatus, error) {
	resp, err := s.Readiness(ctx, &emptypb.Empty{})
	if err != nil {
		return grpc_health_v1.HealthCheckResponse_UNKNOWN, fmt.Errorf("error while checking readiness: %v", err)
	}
	status := grpc_health_v1.HealthCheckResponse_SERVING
	if !resp.Ok {
		status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	}

	return status, nil
}
