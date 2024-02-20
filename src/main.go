package main

import (
	"komeet/cmd"
	. "komeet/config"
	"komeet/core"
)

func init() {
	Config.Load()
	core.InitDatabase()

	log := core.Logger()
	log.Info().Msg("Initializing Komeet app...")
}

func main() {
	cmd.Execute()
}
