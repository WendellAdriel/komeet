package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "komeet/config"
)

var DbEngine *gorm.DB

func InitDatabase() {
	log := Logger()
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		Config.Secrets.DB.User,
		Config.Secrets.DB.Pass,
		Config.Secrets.DB.Host,
		Config.Secrets.DB.Port,
		Config.Secrets.DB.DB,
	)

	DbEngine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msgf("Error opening MySQL %s DB", Config.Secrets.DB.DB)
	}
}
