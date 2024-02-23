package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "komeet/core"
	"komeet/repositories"
	"komeet/utils"
	"net/http"
)

type LoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var dto LoginDTO
	if err := c.ShouldBind(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusUnprocessableEntity, "Invalid data", err)
		return
	}

	token, err := handleLogin(dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnprocessableEntity, "The credentials don't match our records", err)
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
	})
}

func handleLogin(dto LoginDTO) (string, error) {
	user, found := repositories.GetUserForLogin(dto.Email)
	if !found {
		return "", errors.New("invalid login data")
	}

	if user.CheckPassword(dto.Password) == false {
		return "", errors.New("invalid login data")
	}

	return NewToken(user.UUID), nil
}
