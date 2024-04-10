package bot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hemreari/goldwatcher/config"
	"github.com/hemreari/goldwatcher/price"
	"github.com/hemreari/goldwatcher/scrapper"
	log "github.com/sirupsen/logrus"
)

type BotModel interface {
	NewMessageReceived(update tgbotapi.Update) tgbotapi.MessageConfig
	HandleCurrentPricesCmd(chatId int64) tgbotapi.MessageConfig
}

type TgClient struct {
	Conf *config.Config
	Bot  *tgbotapi.BotAPI
	Pm   price.PriceModel
}

func newBotClient(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func NewTgClient(cfg *config.Config, pm price.PriceModel) (*TgClient, error) {
	bot, err := newBotClient(cfg.Tg.Token)
	if err != nil {
		return nil, fmt.Errorf("error while getting tg client: %v", err)
	}

	bot.Debug = cfg.Tg.Debug
	log.Printf("authorized on account %s", bot.Self.UserName)

	return &TgClient{Bot: bot, Pm: pm, Conf: cfg}, nil
}

func (t *TgClient) NewMessageReceived(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if update.Message == nil { // ignore any non-Message updates
		return
	}

	if !update.Message.IsCommand() { // ignore any non-command Messages
		return
	}

	var msg tgbotapi.MessageConfig

	switch update.Message.Command() {
	case "help":
		msg.Text = "I understand /sayhi and /status."
	case "sayhi":
		msg.Text = "Hi :)"
	case "status":
		msg.Text = "I'm ok."
	case "anlik":
		msg = t.HandleCurrentPricesCmd(update.Message.Chat.ID)
	default:
		msg.Text = "I don't know that command"
	}

	if _, err := t.Bot.Send(msg); err != nil {
		log.Errorf("couldn't send the message: %v", err)
	}
}

func (t *TgClient) HandleCurrentPricesCmd(chatId int64) tgbotapi.MessageConfig {
	ctx := context.Background()
	msg := tgbotapi.NewMessage(chatId, "")

	var pri *price.Price
	pri = t.Pm.GetLatestPrice(ctx, t.Conf.App.ExpirationMin)
	if pri == nil {
		pri = scrapper.GetPrices()
		t.Pm.InsertNewPrice(ctx, pri)
	}

	msg.Text = getPriceMsg(*pri)
	return msg
}

func getPriceMsg(price price.Price) string {
	return fmt.Sprintf("22 Ayar Altin:\t\t\t%d\nCeyrek:\t\t\t%d\nYarim:\t\t\t%d\nTam:\t\t\t%d\nCumhuriyet:\t\t%d\n IAB Kapanis:\t%d",
		price.Ayar22Altin, price.Ceyrek, price.Yarim, price.Tam, price.Cumhuriyet, price.IabKapanis,
	)
}
