package core

import (
	"fmt"
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
	App.Router.Use(gin.CustomRecovery(errorHandler))
}

func registerHealthRoute() {
	App.Router.GET("/health", func(c *gin.Context) {
		response := gin.H{
			"status": "OK",
		}

		if !App.IsLocal() {
			c.JSON(http.StatusOK, response)
			return
		}

		response["name"] = App.Config.Name
		response["env"] = App.Config.Env
		response["version"] = runtime.Version()
		response["timestamp"] = time.Now().Format(time.RFC1123)

		c.JSON(http.StatusOK, response)
	})
}

func registerFallbackRoute() {
	App.Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
		})
	})
}

func errorHandler(c *gin.Context, err any) {
	logger := App.Logger()
	logger.Error().Msgf("Unknown Error: %v", err)

	response := gin.H{
		"message": "Internal server error",
	}

	if !App.IsProduction() {
		response["error"] = fmt.Sprintf("%v", err)
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, response)
}
