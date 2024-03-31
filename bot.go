package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgStruct struct {
	Bot *tgbotapi.BotAPI
}

func NewTgStruct(cfg *Config) *TgStruct {
	bot, err := tgbotapi.NewBotAPI(cfg.Tg.Token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = cfg.Tg.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TgStruct{Bot: bot}
}
