package common

import (
	"github.com/gin-gonic/gin"
	. "komeet/config"
	"net/http"
	"runtime"
	"time"
)

func Health(c *gin.Context) {
	if Config.IsLocal() {
		c.JSON(http.StatusOK, gin.H{
			"status":    "OK",
			"name":      Config.Name,
			"env":       Config.Env,
			"version":   runtime.Version(),
			"timestamp": time.Now().Format(time.RFC1123),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
