package cmd

import (
	"github.com/spf13/cobra"
	. "komeet/core"
)

var rootCmd = &cobra.Command{
	Use:   "komeet",
	Short: "Komeet app runner",
}

func Execute() {
	logger := App.Logger()
	if err := rootCmd.Execute(); err != nil {
		logger.Panic().Err(err).Msg("root execute")
	}
}
