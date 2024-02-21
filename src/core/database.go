package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
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
		log.Fatal().Msgf("Error opening MySQL %s DB", App.Secrets.DB.Name)
	}
}
