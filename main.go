package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func readConfig() *Config {
	cfg := &Config{}

	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(cfg); err != nil {
		log.Fatal("Config error: ", err.Error())
	}

	return cfg
}

func main() {
	cfg := readConfig()
	dbClient := NewDbClient(cfg)

	tgClient := NewTgStruct(cfg)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := tgClient.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			tgClient.NewMessageReceived(update)
		}
	}

	price := GetPrices()
	dbClient.InsertNewPrice(&price)
}
