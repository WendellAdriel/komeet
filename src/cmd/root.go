package cmd

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	. "komeet/config"
	"komeet/core"
	"komeet/endpoints/common"
	"net/http"
)

var rootCmd = &cobra.Command{
	Use:   "komeet",
	Short: "Komeet app runner",
}

func Execute() {
	log := core.Logger()

	apiMode := gin.DebugMode
	if Config.IsProduction() {
		apiMode = gin.ReleaseMode
	}

	log.Printf("Starting the API in port %d for %s mode", Config.ApiPort, apiMode)

	gin.SetMode(apiMode)
	router := gin.Default()

	// Serve Vue SPA
	spaLocation := "./dist/web"
	if Config.IsProduction() {
		spaLocation = "./web" // For Production, only the dist folder should be used
	}
	router.Use(static.Serve("/", static.LocalFile(spaLocation, false)))

	// Global Middlewares
	router.Use(core.LoggerMiddleware)

	// Health Route
	router.GET("/health", common.Health)

	// Application Routes

	// Fallback API Route
	router.Any("/api/:any", common.NotFound)

	// Fallback Route for Vue to handle
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})

	router.Run(fmt.Sprintf(`:%d`, Config.ApiPort))
}
