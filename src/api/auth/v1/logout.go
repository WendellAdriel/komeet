package v1

import (
	"github.com/gin-gonic/gin"
	"komeet/repositories"
	"komeet/utils"
)

func Logout(c *gin.Context) {
	user, err := utils.GetUserFromToken(c)
	if err != nil {
		utils.UnauthorizedResponse(c)
		return
	}

	user.AuthUUID = ""
	repositories.UpdateUser(&user)
	utils.NoContentResponse(c)
}
