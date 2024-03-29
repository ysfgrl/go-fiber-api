package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserPayload struct {
	UserID   string `json:"id"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}

type tokenPayload struct {
	UserPayload
	IssuedAt  time.Time `json:"iat,omitempty"`
	ExpiredAt time.Time `json:"exp,omitempty"`
	NotBefore time.Time `json:"nbf,omitempty"`
}

func (t tokenPayload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(t.ExpiredAt), nil
}

func (t tokenPayload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(t.IssuedAt), nil
}

func (t tokenPayload) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(t.NotBefore), nil
}

func (t tokenPayload) GetIssuer() (string, error) {
	return t.UserName, nil
}

func (t tokenPayload) GetSubject() (string, error) {
	return t.UserName, nil
}

func (t tokenPayload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}
