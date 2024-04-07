package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hemreari/goldwatcher/bot"
	"github.com/hemreari/goldwatcher/config"
	"github.com/hemreari/goldwatcher/price"
	log "github.com/sirupsen/logrus"
)

func readConfig() (*config.Config, error) {
	// cfg := &Config{}

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	return nil, err
	// }

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

	cfg := &config.Config{
		Db: config.DbConf{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     port,
		},
		Tg: config.TgConf{
			Token: os.Getenv("TG_TOKEN"),
			Debug: tgDebug,
		},
		App: config.AppConf{
			ExpirationMin: expMin,
		},
	}

	return cfg, nil
}

func main() {
	ctx := context.Background()
	cfg, err := readConfig()
	if err != nil {
		log.Fatalf("couldn't read the env variables: %v", err)
	}
	dbClient, err := price.NewDbClient(ctx, cfg)
	if err != nil {
		log.Fatalf("couldn't create new db client: %v", err)
	}

	tgClient := bot.NewTgClient(cfg, dbClient)

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
