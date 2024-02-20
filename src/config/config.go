package config

import (
	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

var Config Settings

func (s *Settings) Load() {
	var err error
	{
		if b, err := os.ReadFile("config.json"); err == nil {
			b = []byte(strings.TrimSpace(string(b)))
			err = sonic.Unmarshal(b, &Config)
		}
		if err != nil {
			log.Fatal().Err(err).Msg("Loading Config")
		}
	}
	{
		if b, err := os.ReadFile("secrets.json"); err == nil {
			b = []byte(strings.TrimSpace(string(b)))
			err = sonic.Unmarshal(b, &Config.Secrets)
		}
		if err != nil {
			log.Fatal().Err(err).Msg("Loading Secrets")
		}
	}
}

func (s *Settings) IsProduction() bool {
	return strings.ToLower(s.Env) == "production"
}

func (s *Settings) IsLocal() bool {
	return strings.ToLower(s.Env) == "local"
}
