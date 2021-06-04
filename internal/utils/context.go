package utils

import (
	"context"
	"fmt"
	"github.com/peppys/service-template/internal/entities"
)

const (
	UserClaimsKey = "user-claims"
)

func ContextWithUserClaims(ctx context.Context, claims *entities.UserClaims) context.Context {
	return context.WithValue(ctx, UserClaimsKey, claims)
}

func UserClaimsFromContext(ctx context.Context) (*entities.UserClaims, error) {
	u, ok := ctx.Value(UserClaimsKey).(*entities.UserClaims)
	if !ok {
		return nil, fmt.Errorf("user claims not found in context")
	}

	return u, nil
}
