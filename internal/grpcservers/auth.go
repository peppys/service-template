package grpcservers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/peppys/service-template/gen/go/proto"
	"github.com/peppys/service-template/internal/entities"
	"github.com/peppys/service-template/internal/services"
	"github.com/peppys/service-template/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthGrpcServer struct {
	proto.UnimplementedAuthServiceServer
	service authService
}

type authService interface {
	CreateUser(context.Context, *entities.User) (*entities.User, error)
	GenerateTokensViaCredentials(context.Context, string, string) (*services.AuthTokens, error)
	GenerateTokensViaRefreshToken(context.Context, string) (*services.AuthTokens, error)
	GenerateTokens(context.Context, uuid.UUID) (*services.AuthTokens, error)
}

func NewAuthGrpcServer(service *services.AuthService) *AuthGrpcServer {
	return &AuthGrpcServer{service: service}
}

func (a *AuthGrpcServer) Me(ctx context.Context, empty *emptypb.Empty) (*proto.User, error) {
	u, err := utils.UserClaimsFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	return &proto.User{
		Id:         u.Id,
		Email:      u.Email,
		Username:   u.Username,
		GivenName:  u.GivenName,
		FamilyName: u.FamilyName,
		Nickname:   u.GivenName,
		Picture:    u.Picture,
	}, nil
}

func (a *AuthGrpcServer) Signup(ctx context.Context, request *proto.SignupRequest) (*proto.TokenResponse, error) {
	_, err := utils.UserClaimsFromContext(ctx)
	if err == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "you are already authenticated")
	}

	// create user
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("invalid password: %v", err)
	}
	user, err := a.service.CreateUser(ctx, &entities.User{
		Email:        request.GetEmail(),
		Username:     request.GetUsername(),
		PasswordHash: string(passwordHash),
		GivenName:    request.GetGivenName(),
		FamilyName:   request.GetFamilyName(),
		Name:         fmt.Sprintf("%s %s", request.GetGivenName(), request.GetFamilyName()),
		Nickname:     request.GetNickname(),
		Picture:      request.GetPicture(),
	})
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	tokens, err := a.service.GenerateTokens(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating token after signup: %v", err)
	}

	return a.toProto(tokens), nil
}

func (a *AuthGrpcServer) Token(ctx context.Context, request *proto.TokenRequest) (*proto.TokenResponse, error) {
	_, err := utils.UserClaimsFromContext(ctx)
	if err == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "you are already authenticated")
	}

	switch request.GetGrantType() {
	case proto.GrantType_password:
		tokens, err := a.service.GenerateTokensViaCredentials(ctx, request.GetUsername(), request.GetPassword())
		if err != nil {
			return nil, fmt.Errorf("error generating tokens via credentials: %v", err)
		}
		return a.toProto(tokens), nil
	case proto.GrantType_refresh_token:
		tokens, err := a.service.GenerateTokensViaRefreshToken(ctx, request.GetRefreshToken())
		if err != nil {
			return nil, fmt.Errorf("error generating tokens via refresh token: %v", err)
		}
		return a.toProto(tokens), nil
	default:
		return nil, fmt.Errorf("unsupported grant type: %s", request.GetGrantType())
	}
}

func (a *AuthGrpcServer) toProto(tokens *services.AuthTokens) *proto.TokenResponse {
	return &proto.TokenResponse{
		AccessToken:      tokens.AccessToken,
		TokenType:        "bearer",
		ExpiresIn:        int32(tokens.AccessTokenExpiresIn.Seconds()),
		RefreshToken:     tokens.RefreshToken,
		RefreshExpiresIn: int32(tokens.RefreshTokenExpiresIn.Seconds()),
	}
}
