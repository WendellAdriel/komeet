package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "komeet/core"
	"komeet/models"
	"komeet/repositories"
)

func GetUserFromToken(c *gin.Context) (models.User, error) {
	token := c.GetHeader(App.Config.Auth.TokenHeader)
	tokenClaims, err := ParseTokenClaims(token)
	if err != nil {
		UnauthorizedResponse(c)
		return models.User{}, err
	}

	user, found := repositories.GetUserBy("uuid", tokenClaims.UserUUID)
	if !found {
		UnauthorizedResponse(c)
		return models.User{}, errors.New("unauthorized")
	}

	if user.AuthUUID != tokenClaims.AuthUUID {
		UnauthorizedResponse(c)
		return models.User{}, errors.New("unauthorized")
	}

	return user, nil
}
