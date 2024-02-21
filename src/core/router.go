package core

import (
	"github.com/gin-contrib/static"
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
	registerSPARoute()
	registerHealthRoute()
	registerFallbackRoutes()
}

func registerGlobalMiddlewares() {
	App.Router.Use(loggerMiddleware)
}

func registerSPARoute() {
	spaLocation := "./dist/web"
	if !App.IsLocal() {
		spaLocation = "./web" // For Non-local envs, only the dist folder should be used
	}
	App.Router.Use(static.Serve("/", static.LocalFile(spaLocation, false)))
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

func registerFallbackRoutes() {
	// API fallback route
	App.Router.Any("/api/:any", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
		})
	})

	// SPA fallback route
	App.Router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
}
