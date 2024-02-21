package core

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"komeet/config"
	"os"
	"strings"
)

type Application struct {
	Config  config.Settings
	Secrets config.Secrets
	DB      *gorm.DB
	Router  *gin.Engine
}

var App Application

func (a *Application) IsProduction() bool {
	return strings.ToLower(a.Config.Env) == "production"
}

func (a *Application) IsLocal() bool {
	return strings.ToLower(a.Config.Env) == "local"
}

func (a *Application) Init() {
	logger := a.Logger()
	logger.Info().Msg("Initializing Komeet app...")

	loadConfigs()
	initDatabase()
	migrateDatabase()
	initRouter()
}

func (a *Application) Run() {
	a.Router.Run(fmt.Sprintf(`:%d`, a.Config.Port))
}

func loadConfigs() {
	logger := App.Logger()
	var err error
	{
		if b, err := os.ReadFile("config.json"); err == nil {
			b = []byte(strings.TrimSpace(string(b)))
			err = sonic.Unmarshal(b, &App.Config)
		}
		if err != nil {
			logger.Fatal().Err(err).Msg("Loading Config")
		}
	}
	{
		if b, err := os.ReadFile("secrets.json"); err == nil {
			b = []byte(strings.TrimSpace(string(b)))
			err = sonic.Unmarshal(b, &App.Secrets)
		}
		if err != nil {
			logger.Fatal().Err(err).Msg("Loading Secrets")
		}
	}
}
