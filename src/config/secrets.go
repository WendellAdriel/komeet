package config

type Secrets struct {
	DB DBCreds `json:"db"`
}

type DBCreds struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	DB   string `json:"db"`
	User string `json:"user"`
	Pass string `json:"pass"`
}
