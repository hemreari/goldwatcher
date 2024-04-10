package bot

import (
	"context"
	"fmt"
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hemreari/goldwatcher/config"
	"github.com/hemreari/goldwatcher/price"
)

type MockPriceModel struct{}

func (m *MockPriceModel) InsertNewPrice(_ context.Context, _ *price.Price) {
}

func (m *MockPriceModel) GetLatestPrice(_ context.Context, expirationMin int) *price.Price {
	return &price.Price{
		Id:          1,
		LastAt:      time.Now(),
		Ayar22Altin: 1111,
		Ceyrek:      2222,
		Yarim:       3333,
		Tam:         4444,
		Cumhuriyet:  5555,
		IabKapanis:  6666,
	}
}

type MockBotModel struct {
	mockPriceModel *MockPriceModel
}

func (m *MockBotModel) newBotClient(_ string) (*tgbotapi.BotAPI, error) {
	return &tgbotapi.BotAPI{Self: tgbotapi.User{UserName: "test_api"}}, nil
}

func TestNewTgClient(t *testing.T) {
	cfg := &config.Config{}
	tgClient, err := NewTgClient(cfg, &MockPriceModel{})
	if tgClient == nil {
		t.Error("received tg client nil value")
	}

	if err != nil {
		t.Error("received err non nil value.")
	}
}

func TestHandleCurrentPricesCmd(t *testing.T) {
	tgClient, _ := NewTgClient(&config.Config{}, &MockPriceModel{})

	tgClient.HandleCurrentPricesCmd(123)
}

func TestGetPriceMsg(t *testing.T) {
	lastAt := time.Now()
	priceStruct := price.Price{
		Id:          1,
		LastAt:      lastAt,
		Ayar22Altin: 1111,
		Ceyrek:      2222,
		Yarim:       3333,
		Tam:         4444,
		Cumhuriyet:  5555,
		IabKapanis:  6666,
	}

	msg := getPriceMsg(priceStruct)
	expectedMsg := fmt.Sprintf("22 Ayar Altin:\t\t\t%d\nCeyrek:\t\t\t%d\nYarim:\t\t\t%d\nTam:\t\t\t%d\nCumhuriyet:\t\t%d\n IAB Kapanis:\t%d",
		priceStruct.Ayar22Altin,
		priceStruct.Ceyrek,
		priceStruct.Yarim,
		priceStruct.Tam,
		priceStruct.Cumhuriyet,
		priceStruct.IabKapanis,
	)

	if msg != expectedMsg {
		t.Errorf("price msg is not equal to expected price msg")
	}
}
