package core

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type UserClaims struct {
	UserUUID string `json:"user_uuid"`
	AuthUUID string `json:"auth_uuid"`
	jwt.RegisteredClaims
}

func NewToken(userUuid string) (string, string) {
	authUuid := uuid.NewString()
	claims := UserClaims{
		userUuid,
		authUuid,
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

	return signed, authUuid
}

func ParseTokenClaims(token string) (UserClaims, error) {
	var userClaims UserClaims

	jwtToken, err := jwt.ParseWithClaims(token, &userClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(App.Secrets.JWTKey), nil
	})

	if err != nil {
		return userClaims, err
	}

	if !jwtToken.Valid {
		return userClaims, errors.New("invalid token")
	}

	return userClaims, nil
}
