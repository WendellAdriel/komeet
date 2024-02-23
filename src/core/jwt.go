package core

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

func NewToken(userUuid string) string {
	claims := UserClaims{
		userUuid,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * App.Config.Auth.TokenValidity)),
		},
	}

	logger := App.Logger()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(App.Secrets.JWTKey))
	if err != nil {
		logger.Panic().Err(err).Msg("Failed to generate JWT")
	}

	return signed
}
