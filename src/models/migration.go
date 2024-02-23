package models

import (
	. "komeet/core"
)

func Migrate() {
	logger := App.Logger()
	err := App.DB.AutoMigrate(&User{})
	if err != nil {
		logger.Panic().Err(err).Msgf("Error migrating MySQL %s DB", App.Secrets.DB.Name)
	}
}
