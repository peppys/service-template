package interceptors

import (
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/peppys/service-template/internal/services"
	"github.com/peppys/service-template/internal/utils"
	"google.golang.org/grpc"
	"log"
)

func Authorization(authService *services.AuthService) grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return ctx, nil
		}

		claims, err := authService.VerifyToken(ctx, token)
		if err != nil {
			log.Printf("Failed verifying token: %v", token)
			return ctx, nil
		}

		log.Printf("Attatching auth to request: %v", claims)
		ctx = utils.ContextWithUserClaims(ctx, claims)
		grpc_ctxtags.Extract(ctx).Set("auth.sub", claims.Id)

		return ctx, nil
	})
}
