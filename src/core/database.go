package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	logger := App.Logger()
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
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
