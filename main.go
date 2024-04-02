package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

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
	ctx := context.Background()
	cfg := readConfig()
	dbClient := NewDbClient(ctx, cfg)

	tgClient := NewTgClient(cfg, dbClient)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := tgClient.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			tgClient.NewMessageReceived(update)
		}
	}

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)
	<-gracefulShutdown
	shutdown(ctx, dbClient)
	fmt.Println("Done!")
}

func shutdown(ctx context.Context, dbClient *DbClient) {
	log.Print("shutting down the application...")
	dbClient.Db.Close(ctx)
}
