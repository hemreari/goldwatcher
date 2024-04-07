package config

type Config struct {
	Tg  TgConf  `json:"telegram"`
	Db  DbConf  `json:"db"`
	App AppConf `json:"app"`
}

type DbConf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
	Port     int    `json:"port"`
}

type TgConf struct {
	Token string `json:"token"`
	Debug bool   `json:"debug"`
}

type AppConf struct {
	ExpirationMin int `json:"expirationMin"`
}
