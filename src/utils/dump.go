package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DumpAbort(c *gin.Context, value any) {
	c.AbortWithStatusJSON(http.StatusTeapot, value)
}
