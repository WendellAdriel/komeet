package main

import (
	"komeet/cmd"
	. "komeet/core"
	"komeet/models"
)

func init() {
	App.Init()
	models.Migrate()
}

func main() {
	cmd.Execute()
}
