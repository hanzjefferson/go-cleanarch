package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type Provider struct {
	Issuer string
	Secret []byte
	ExpiredDuration time.Duration
}

func (jwtProvider *Provider) Generate(userId uint) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: jwtProvider.Issuer,
			IssuedAt: jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtProvider.ExpiredDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtProvider.Secret)
}