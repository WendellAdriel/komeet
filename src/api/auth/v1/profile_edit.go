package v1

import (
	"github.com/gin-gonic/gin"
	"komeet/models"
	"komeet/utils"
	"net/http"
	"time"
)

type EditProfileDTO struct {
	Name                 string `json:"name"`
	CurrentPassword      string `json:"current_password"`
	NewPassword          string `json:"new_password" binding:"required_with=current_password"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required_with=new_password"`
}

func EditProfile(c *gin.Context) {
	user, err := utils.GetUserFromToken(c)
	if err != nil {
		utils.UnauthorizedResponse(c)
		return
	}

	var dto EditProfileDTO
	if err := c.ShouldBind(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusUnprocessableEntity, "Invalid data", err)
		return
	}

	user, err = handleEditProfile(user, dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update profile", err)
		return
	}

	c.JSON(http.StatusOK, ProfileResponse{
		UUID:      user.UUID,
		Name:      user.Name,
		Email:     user.Email,
		Active:    user.Active,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	})
}

func handleEditProfile(user models.User, dto EditProfileDTO) (models.User, error) {
	// TODO: Implement
	return user, nil
}
