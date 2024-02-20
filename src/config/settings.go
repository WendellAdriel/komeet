package config

type Settings struct {
	Env     string  `json:"env"`
	Name    string  `json:"name"`
	ApiPort int     `json:"api_port"`
	Logs    Logs    `json:"logs"`
	Secrets Secrets `json:"-"`
}

type Logs struct {
	MaxBackups int  `json:"max_backups"`
	MaxSize    int  `json:"max_size"`
	MaxDays    int  `json:"max_days"`
	Compress   bool `json:"compress"`
}
