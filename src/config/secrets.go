package config

type Secrets struct {
	JWTKey string  `json:"jwt_key"`
	DB     DBCreds `json:"db"`
}

type DBCreds struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Name string `json:"name"`
	User string `json:"user"`
	Pass string `json:"pass"`
}
