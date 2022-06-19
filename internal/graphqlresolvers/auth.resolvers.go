package graphqlresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graphql1 "github.com/peppys/service-template/gen/go/graphql"
	"github.com/peppys/service-template/internal/entities"
	"github.com/peppys/service-template/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) AuthTokens(ctx context.Context, input *graphql1.TokensInput) (*graphql1.Tokens, error) {
	u, err := utils.UserClaimsFromContext(ctx)
	if err == nil {
		return nil, fmt.Errorf("you are already authenticated as %s", u.Email)
	}

	switch input.GrantType {
	case graphql1.GrantTypePassword:
		if input.Username == nil {
			return nil, fmt.Errorf("username is required for grant type: %s", input.GrantType)
		}
		if input.Password == nil {
			return nil, fmt.Errorf("password is required for grant type: %s", input.GrantType)
		}
		tokens, err := r.authService.GenerateTokensViaCredentials(ctx, *input.Username, *input.Password)
		if err != nil {
			return nil, fmt.Errorf("error generating tokens via credentials: %v", err)
		}
		return &graphql1.Tokens{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
			TokenType:    graphql1.TokenTypeBearer,
			Expires:      tokens.AccessTokenExpiresIn.Seconds(),
		}, nil
	case graphql1.GrantTypeRefreshToken:
		if input.RefreshToken == nil {
			return nil, fmt.Errorf("refresh token is required for grant type: %s", input.GrantType)
		}
		tokens, err := r.authService.GenerateTokensViaRefreshToken(ctx, *input.RefreshToken)
		if err != nil {
			return nil, fmt.Errorf("error generating tokens via refresh token: %v", err)
		}
		return &graphql1.Tokens{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
			TokenType:    graphql1.TokenTypeBearer,
			Expires:      tokens.AccessTokenExpiresIn.Seconds(),
		}, nil
	default:
		return nil, fmt.Errorf("unsupported grant type: %s", input.GrantType)
	}
}

func (r *mutationResolver) Signup(ctx context.Context, input *graphql1.SignupInput) (*graphql1.Tokens, error) {
	u, err := utils.UserClaimsFromContext(ctx)
	if err == nil {
		return nil, fmt.Errorf("you are already authenticated as %s", u.Email)
	}

	// create user
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("invalid password: %v", err)
	}
	user, err := r.authService.CreateUser(ctx, &entities.User{
		Email:        input.Email,
		Username:     input.Username,
		PasswordHash: string(passwordHash),
		GivenName:    input.GivenName,
		FamilyName:   input.FamilyName,
		Name:         fmt.Sprintf("%s %s", input.GivenName, input.FamilyName),
		Nickname:     input.Nickname,
		Picture:      input.Picture,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	tokens, err := r.authService.GenerateTokens(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating token after signup: %v", err)
	}

	return &graphql1.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    graphql1.TokenTypeBearer,
		Expires:      tokens.AccessTokenExpiresIn.Seconds(),
	}, nil
}

func (r *queryResolver) WhoAmI(ctx context.Context) (*graphql1.User, error) {
	u, err := utils.UserClaimsFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	return &graphql1.User{
		ID:         u.Id,
		Email:      u.Email,
		Username:   u.Username,
		GivenName:  u.GivenName,
		FamilyName: u.FamilyName,
		Nickname:   u.Nickname,
		Picture:    u.Picture,
	}, nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
