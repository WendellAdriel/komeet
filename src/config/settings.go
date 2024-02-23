package config

import "time"

type Settings struct {
	Env  string       `json:"app_env"`
	Name string       `json:"app_name"`
	Port int          `json:"app_port"`
	Auth AuthSettings `json:"auth"`
	Logs Logs         `json:"logs"`
}

type AuthSettings struct {
	TokenHeader   string        `json:"token_header"`
	TokenValidity time.Duration `json:"token_validity"`
}

type Logs struct {
	MaxBackups int  `json:"max_backups"`
	MaxSize    int  `json:"max_size"`
	MaxDays    int  `json:"max_days"`
	Compress   bool `json:"compress"`
}
