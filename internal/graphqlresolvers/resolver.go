package graphqlresolvers

import (
	"context"
	"github.com/google/uuid"
	"github.com/peppys/service-template/internal/entities"
	"github.com/peppys/service-template/internal/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type authService interface {
	CreateUser(context.Context, *entities.User) (*entities.User, error)
	GenerateTokensViaCredentials(context.Context, string, string) (*services.AuthTokens, error)
	GenerateTokensViaRefreshToken(context.Context, string) (*services.AuthTokens, error)
	GenerateTokens(context.Context, uuid.UUID) (*services.AuthTokens, error)
}

type Resolver struct {
	authService
}

func New(authService authService) *Resolver {
	return &Resolver{
		authService,
	}
}
