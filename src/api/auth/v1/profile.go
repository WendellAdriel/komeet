package v1

import (
	"github.com/gin-gonic/gin"
	"komeet/utils"
	"net/http"
	"time"
)

type ProfileResponse struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Profile(c *gin.Context) {
	user, err := utils.GetUserFromToken(c)
	if err != nil {
		utils.UnauthorizedResponse(c)
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
