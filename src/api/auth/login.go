package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "komeet/core"
	"komeet/models/auth"
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

	token, err := handleLogin(dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func handleLogin(dto LoginDTO) (string, error) {
	var user auth.User

	App.DB.Where("email = ?", dto.Email).
		Where("email_verified_at IS NOT NULL").
		Where("active = ?", true).
		First(&user)

	if user.ID == 0 {
		return "", errors.New("invalid login data")
	}

	if user.CheckPassword(dto.Password) == false {
		return "", errors.New("invalid login data")
	}

	return NewToken(user.UUID), nil
}
