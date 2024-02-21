package models

import (
	. "komeet/core"
	"komeet/models/auth"
)

func Migrate() {
	logger := App.Logger()
	err := App.DB.AutoMigrate(&auth.User{})
	if err != nil {
		logger.Fatal().Err(err).Msgf("Error migrating MySQL %s DB", App.Secrets.DB.Name)
	}
}
