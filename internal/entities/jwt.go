package entities

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type UserClaims struct {
	jwt.StandardClaims
	UUID       uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	GivenName  string    `json:"given_name"`
	FamilyName string    `json:"family_name"`
	Nickname   string    `json:"nickname"`
	Picture    string    `json:"picture"`
}
