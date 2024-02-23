package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileResponse struct {
}

func Profile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented",
	})
}
