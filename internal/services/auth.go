package services

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/peppys/service-template/internal/entities"
	"github.com/peppys/service-template/internal/repositories"
	"github.com/peppys/service-template/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	user         *repositories.UserRepository
	refreshToken *repositories.RefreshTokenRepository
}

type UserClaims struct {
	jwt.StandardClaims
	Email      string `json:"email"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Nickname   string `json:"nickname"`
	Picture    string `json:"picture"`
}

type RefreshTokenWithSecret struct {
	Secret string
	*entities.RefreshToken
}

type AuthTokens struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresIn  time.Duration
	RefreshTokenExpiresIn time.Duration
}

func NewAuthService(user *repositories.UserRepository, refreshToken *repositories.RefreshTokenRepository) *AuthService {
	return &AuthService{user, refreshToken}
}

func (s *AuthService) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	return s.user.Save(ctx, user)
}

func (s *AuthService) VerifyToken(ctx context.Context, accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte("some-key"), nil
		})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (s *AuthService) GenerateTokensViaCredentials(ctx context.Context, username string, password string) (*AuthTokens, error) {
	user, err := s.user.FindFirst(ctx, entities.User{
		Username: username,
	})
	if err != nil {
		return nil, fmt.Errorf("error finding user with username %s: %v", username, err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, fmt.Errorf("incorrect username password combo: %v", err)
	}
	tokens, err := s.generateTokens(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("error generating new refresh token: %v", err)
	}
	return tokens, nil
}

func (s *AuthService) GenerateTokensViaRefreshToken(ctx context.Context, token string) (*AuthTokens, error) {
	record, err := s.refreshToken.FindFirst(ctx, entities.RefreshToken{
		TokenHash: utils.Md5Hash(token),
	})
	if err != nil {
		return nil, fmt.Errorf("error finding refresh token: %v", err)
	}
	tokens, err := s.generateTokens(ctx, record.UserID)
	if err != nil {
		return nil, fmt.Errorf("error generating new refresh token: %v", err)
	}
	return tokens, nil
}

func (s *AuthService) generateTokens(ctx context.Context, userID uuid.UUID) (*AuthTokens, error) {
	user, err := s.user.FindFirst(ctx, entities.User{
		ID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("error finding user ID %s: %v", userID.String(), err)
	}

	// create new refresh token
	refreshToken, err := s.generateRefreshToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error generating refresh token: %v", err)
	}
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %v", err)
	}

	return &AuthTokens{
		AccessToken:           accessToken,
		AccessTokenExpiresIn:  time.Hour,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: time.Hour * 24 * 7,
	}, nil
}

func (s *AuthService) generateAccessToken(user *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Email:      user.Email,
		Username:   user.Username,
		Name:       user.Name,
		GivenName:  user.GivenName,
		FamilyName: user.FamilyName,
		Nickname:   user.Nickname,
		Picture:    user.Picture,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.ID.String(),
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "api",
		},
	})
	signed, err := token.SignedString([]byte("some-key"))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signed, nil
}

func (s *AuthService) generateRefreshToken(ctx context.Context, user *entities.User) (string, error) {
	// invalidate any existing refresh tokens
	err := s.refreshToken.Delete(ctx, entities.RefreshToken{
		UserID: user.ID,
	})
	if err != nil {
		return "", fmt.Errorf("error invalidating refresh tokens: %v", err)
	}

	refreshToken := b64.StdEncoding.EncodeToString([]byte(uuid.New().String()))
	_, err = s.refreshToken.Save(ctx, &entities.RefreshToken{
		UserID:    user.ID,
		TokenHash: utils.Md5Hash(refreshToken),
		ExpiresAt: time.Now().AddDate(0, 0, 7),
	})
	if err != nil {
		return "", fmt.Errorf("error saving refresh token: %v", err)
	}

	return refreshToken, nil
}
