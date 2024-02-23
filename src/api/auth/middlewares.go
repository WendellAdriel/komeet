package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "komeet/core"
	"komeet/models"
	"komeet/repositories"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	_, err := GetUserFromToken(c)
	if err != nil {
		unauthorizedResponse(c)
		return
	}
	c.Next()
}

func GetUserFromToken(c *gin.Context) (models.User, error) {
	token := c.GetHeader(App.Config.Auth.TokenHeader)
	tokenClaims, err := ParseTokenClaims(token)
	if err != nil {
		unauthorizedResponse(c)
		return models.User{}, err
	}

	user, found := repositories.GetUserBy("uuid", tokenClaims.UserUUID)
	if !found {
		unauthorizedResponse(c)
		return models.User{}, errors.New("unauthorized")
	}

	if user.AuthUUID != tokenClaims.AuthUUID {
		unauthorizedResponse(c)
		return models.User{}, errors.New("unauthorized")
	}

	return user, nil
}

func unauthorizedResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
}
