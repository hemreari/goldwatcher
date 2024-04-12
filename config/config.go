package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

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

func ReadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Errorf("error while loading env values from .env file: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("couldn't convert the DB_PORT env variable to int: %v", err)
	}

	tgDebug, err := strconv.ParseBool(os.Getenv("TG_DEBUG"))
	if err != nil {
		log.Errorf("couldn't convert the TG_DEBUG env variable to bool: %v", err)
		tgDebug = false
	}

	expMin, err := strconv.Atoi(os.Getenv("PRICE_EXPIRATION_MIN"))
	if err != nil {
		log.Errorf("couldn't convert the PRICE_EXPIRATION_MIN env variable to int: %v", err)
		expMin = 5
	}

	cfg := &Config{
		Db: DbConf{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     port,
		},
		Tg: TgConf{
			Token: os.Getenv("TG_TOKEN"),
			Debug: tgDebug,
		},
		App: AppConf{
			ExpirationMin: expMin,
		},
	}

	return cfg, nil
}
