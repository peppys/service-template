package grpcservers

import (
	"context"
	"fmt"
	"github.com/peppys/service-template/gen/go/proto"
	"github.com/peppys/service-template/internal/entities"
	"github.com/peppys/service-template/internal/services"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthGrpcServer struct {
	proto.UnimplementedAuthServiceServer
	service *services.AuthService
}

func NewAuthGrpcServer(service *services.AuthService) *AuthGrpcServer {
	return &AuthGrpcServer{service: service}
}

func (a *AuthGrpcServer) Me(ctx context.Context, empty *emptypb.Empty) (*proto.User, error) {
	u, ok := ctx.Value("user").(*services.UserClaims)
	if !ok {
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
	tokens, err := a.service.GenerateTokensViaCredentials(ctx, user.Username, request.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("error creating token after signup: %v", err)
	}

	return a.toProto(tokens), nil
}

func (a *AuthGrpcServer) Token(ctx context.Context, request *proto.TokenRequest) (*proto.TokenResponse, error) {
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
