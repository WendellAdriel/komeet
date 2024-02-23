package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EditProfileDTO struct {
}

func EditProfile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented",
	})
}
