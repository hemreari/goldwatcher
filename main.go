package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hemreari/goldwatcher/bot"
	"github.com/hemreari/goldwatcher/config"
	"github.com/hemreari/goldwatcher/price"
	"github.com/hemreari/goldwatcher/scrapper"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("couldn't read the env variables: %v", err)
	}
	dbClient, err := price.NewDbClient(ctx, cfg)
	if err != nil {
		log.Fatalf("couldn't create new db client: %v", err)
	}

	scrapperClient := scrapper.NewScrapperClient()

	botclientStruct := bot.NewBotClientStruct()

	tgClient, err := bot.NewTgClient(botclientStruct, dbClient, scrapperClient, cfg)
	if err != nil {
		log.Panic(err)
	}

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

func shutdown(ctx context.Context, dbClient *price.DbClient) {
	log.Print("shutting down the application...")
	dbClient.Db.Close(ctx)
}
