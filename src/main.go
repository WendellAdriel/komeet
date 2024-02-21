package main

import (
	"komeet/cmd"
	. "komeet/core"
)

func init() {
	App.Init()
}

func main() {
	cmd.Execute()
}
