// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
)

type SignupInput struct {
	Email      string  `json:"email"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	GivenName  string  `json:"givenName"`
	FamilyName string  `json:"familyName"`
	Nickname   *string `json:"nickname"`
	Picture    *string `json:"picture"`
}

type Tokens struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	TokenType    TokenType `json:"tokenType"`
	Expires      float64   `json:"expires"`
}

type TokensInput struct {
	GrantType    GrantType `json:"grantType"`
	RefreshToken *string   `json:"refreshToken"`
	Username     *string   `json:"username"`
	Password     *string   `json:"password"`
}

type User struct {
	ID         string  `json:"id"`
	Email      string  `json:"email"`
	Username   string  `json:"username"`
	GivenName  string  `json:"givenName"`
	FamilyName string  `json:"familyName"`
	Nickname   *string `json:"nickname"`
	Picture    *string `json:"picture"`
}

type GrantType string

const (
	GrantTypePassword     GrantType = "password"
	GrantTypeRefreshToken GrantType = "refresh_token"
)

var AllGrantType = []GrantType{
	GrantTypePassword,
	GrantTypeRefreshToken,
}

func (e GrantType) IsValid() bool {
	switch e {
	case GrantTypePassword, GrantTypeRefreshToken:
		return true
	}
	return false
}

func (e GrantType) String() string {
	return string(e)
}

func (e *GrantType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GrantType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GrantType", str)
	}
	return nil
}

func (e GrantType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TokenType string

const (
	TokenTypeBearer TokenType = "Bearer"
)

var AllTokenType = []TokenType{
	TokenTypeBearer,
}

func (e TokenType) IsValid() bool {
	switch e {
	case TokenTypeBearer:
		return true
	}
	return false
}

func (e TokenType) String() string {
	return string(e)
}

func (e *TokenType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TokenType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TokenType", str)
	}
	return nil
}

func (e TokenType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
