package interceptors

import (
	"context"
	"github.com/peppys/service-template/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

func Authorization(authService *services.AuthService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}

		values := md["authorization"]
		if len(values) == 0 {
			return handler(ctx, req)
		}

		accessToken := strings.Split(values[0], "Bearer ")[1]
		claims, err := authService.VerifyToken(ctx, accessToken)
		if err != nil {
			return handler(ctx, req)
		}

		log.Printf("Attatching auth to request: %v", claims)
		ctx = context.WithValue(ctx, "user", claims)

		return handler(ctx, req)
	}
}
