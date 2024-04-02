package main

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgClient struct {
	Conf *Config
	Bot  *tgbotapi.BotAPI
	DB   *DbClient
}

func NewTgClient(cfg *Config, dbClient *DbClient) *TgClient {
	bot, err := tgbotapi.NewBotAPI(cfg.Tg.Token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = cfg.Tg.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TgClient{Bot: bot, DB: dbClient, Conf: cfg}
}

func (t *TgClient) NewMessageReceived(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if update.Message == nil { // ignore any non-Message updates
		return
	}

	if !update.Message.IsCommand() { // ignore any non-command Messages
		return
	}

	// Create a new MessageConfig. We don't have text yet,
	// so we leave it empty.

	var msg tgbotapi.MessageConfig
	// Extract the command from the Message.

	switch update.Message.Command() {
	case "help":
		msg.Text = "I understand /sayhi and /status."
	case "sayhi":
		msg.Text = "Hi :)"
	case "status":
		msg.Text = "I'm ok."
	case "anlik":
		msg = t.handleCurrentPricesCmd(update.Message.Chat.ID)
	default:
		msg.Text = "I don't know that command"
	}

	if _, err := t.Bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

/*
22 Ayar Altin: 	2396
Ceyrek:			4420
Yarim:			8840
Tam:			17680
Cumhuriyet:		18002
IAB Kapanis:	2446
*/

func (t *TgClient) handleCurrentPricesCmd(chatId int64) tgbotapi.MessageConfig {
	ctx := context.Background()
	msg := tgbotapi.NewMessage(chatId, "")

	var price *Price
	price = t.DB.GetLatestPrice(ctx, t.Conf.App.ExpirationMin)
	if price == nil {
		price = GetPrices()
		t.DB.InsertNewPrice(ctx, price)
	}

	msg.Text = getPriceMsg(price)
	return msg
}

func getPriceMsg(price *Price) string {
	return fmt.Sprintf("22 Ayar Altin:\t\t\t%d\nCeyrek:\t\t\t%d\nYarim:\t\t\t%d\nTam:\t\t\t%d\nCumhuriyet:\t\t%d\n IAB Kapanis:\t%d",
		price.Ayar22Altin, price.Ceyrek, price.Yarim, price.Tam, price.Cumhuriyet, price.IabKapanis,
	)
}
