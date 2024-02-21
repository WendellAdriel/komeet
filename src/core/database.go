package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"komeet/models/auth"
)

func initDatabase() {
	logger := App.Logger()
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		App.Secrets.DB.User,
		App.Secrets.DB.Pass,
		App.Secrets.DB.Host,
		App.Secrets.DB.Port,
		App.Secrets.DB.Name,
	)

	App.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal().Msgf("Error opening MySQL %s DB", App.Secrets.DB.Name)
	}
}

func migrateDatabase() {
	logger := App.Logger()
	err := App.DB.AutoMigrate(&auth.User{})
	if err != nil {
		logger.Fatal().Err(err).Msgf("Error migrating MySQL %s DB", App.Secrets.DB.Name)
	}
}
