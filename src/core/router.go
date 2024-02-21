package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

func initRouter() {
	apiMode := gin.DebugMode
	if App.IsProduction() {
		apiMode = gin.ReleaseMode
	}

	logger := App.Logger()
	logger.Info().Msgf("Starting the API in port %d for %s mode", App.Config.Port, apiMode)

	gin.SetMode(apiMode)
	App.Router = gin.Default()

	registerGlobalMiddlewares()
	registerHealthRoute()
	registerFallbackRoute()
}

func registerGlobalMiddlewares() {
	App.Router.Use(loggerMiddleware)
}

func registerHealthRoute() {
	App.Router.GET("/health", func(c *gin.Context) {
		if App.IsLocal() {
			c.JSON(http.StatusOK, gin.H{
				"status":    "OK",
				"name":      App.Config.Name,
				"env":       App.Config.Env,
				"version":   runtime.Version(),
				"timestamp": time.Now().Format(time.RFC1123),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}

func registerFallbackRoute() {
	App.Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
		})
	})
}
