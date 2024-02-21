package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func editProfile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented",
	})
}
