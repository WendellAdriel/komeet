package auth

import (
	"github.com/gin-gonic/gin"
	. "komeet/core"
	"komeet/repositories"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	logger := App.Logger()

	token := c.GetHeader(App.Config.Auth.TokenHeader)
	tokenClaims, err := ParseTokenClaims(token)
	if err != nil {
		logger.Error().Err(err).Msg("Token not parsed")
		unauthorizedResponse(c)
		return
	}

	logger.Info().Msgf("token claims %v", tokenClaims)

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
