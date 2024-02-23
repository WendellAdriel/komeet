package auth

import (
	"github.com/gin-gonic/gin"
	. "komeet/core"
	"komeet/repositories"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	token := c.GetHeader(App.Config.Auth.TokenHeader)
	tokenClaims, err := ParseTokenClaims(token)
	if err != nil {
		unauthorizedResponse(c)
		return
	}

	user, found := repositories.GetUserBy("uuid", tokenClaims.UserUUID)
	if !found {
		unauthorizedResponse(c)
		return
	}

	if user.AuthUUID != tokenClaims.AuthUUID {
		unauthorizedResponse(c)
		return
	}

	c.Next()
}

func unauthorizedResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
}
