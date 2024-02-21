package auth

import (
	"github.com/gin-gonic/gin"
	"komeet/utils"
	"net/http"
)

type LoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func login(c *gin.Context) {
	var dto LoginDTO
	if err := c.ShouldBind(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	handleLogin(dto)

	c.JSON(http.StatusNotImplemented, gin.H{
		"email":    dto.Email,
		"password": dto.Password,
	})
}

func handleLogin(dto LoginDTO) {
	// TODO:Implement
}
