package auth

import (
	"github.com/gin-gonic/gin"
	"komeet/utils"
)

func AuthRequired(c *gin.Context) {
	_, err := utils.GetUserFromToken(c)
	if err != nil {
		utils.UnauthorizedResponse(c)
		return
	}
	c.Next()
}
