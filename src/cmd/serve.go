package cmd

import (
	"github.com/spf13/cobra"
	"komeet/api/auth"
	. "komeet/core"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the application server",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	auth.RegisterRoutes()
	// Register your application routes here before running the application
	App.Run()
}
