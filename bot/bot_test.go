package bot

import (
	"context"
	"fmt"
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hemreari/goldwatcher/config"
	"github.com/hemreari/goldwatcher/price"
	"github.com/stretchr/testify/assert"
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

type MockScrapperModel struct{}

func (ms *MockScrapperModel) GetPrices() *price.Price {
	return &price.Price{}
}

func TestNewTgClient(t *testing.T) {
	cfg := &config.Config{}
	tgClient, err := NewTgClient(&MockBotModel{}, &MockPriceModel{}, &MockScrapperModel{}, cfg)

	assert.NotNil(t, tgClient)

	assert.NoError(t, err)
}

func TestHandleCurrentPricesCmd(t *testing.T) {
	tgClient, _ := NewTgClient(&MockBotModel{}, &MockPriceModel{}, &MockScrapperModel{}, &config.Config{})

	expectedMsg := fmt.Sprintf("22 Ayar Altin:\t\t\t%d\nCeyrek:\t\t\t%d\nYarim:\t\t\t%d\nTam:\t\t\t%d\nCumhuriyet:\t\t%d\n IAB Kapanis:\t%d",
		1111, 2222, 3333, 4444, 5555, 6666,
	)

	msgConf := tgClient.HandleCurrentPricesCmd(123)

	assert.Equal(t, msgConf.Text, expectedMsg)

	tgFailedClient, _ := NewTgClient(&MockBotModel{}, &MockFailedPriceModel{}, &MockFailedScrapperModel{}, &config.Config{})
	msgConfFailed := tgFailedClient.HandleCurrentPricesCmd(123)
	expectedFailedMsg := "Couldn't get the latest prices."
	assert.Equal(t, msgConfFailed.Text, expectedFailedMsg)
}

type MockFailedPriceModel struct{}

func (mf *MockFailedPriceModel) GetLatestPrice(_ context.Context, _ int) *price.Price {
	return nil
}

func (mf *MockFailedPriceModel) InsertNewPrice(ctx context.Context, price *price.Price) {
}

type MockFailedScrapperModel struct{}

func (mfs *MockFailedScrapperModel) GetPrices() *price.Price {
	return nil
}

func TestGetPrices(t *testing.T) {
	ctx := context.Background()

	tgClient, _ := NewTgClient(&MockBotModel{}, &MockPriceModel{}, &MockScrapperModel{}, &config.Config{})
	price := tgClient.GetPrices(ctx)

	assert.NotNil(t, price)

	tgFailedClient, _ := NewTgClient(&MockBotModel{}, &MockFailedPriceModel{}, &MockFailedScrapperModel{}, &config.Config{})
	priceFailed := tgFailedClient.GetPrices(ctx)

	assert.Nil(t, priceFailed)
}

func TestNewMessageReceived(t *testing.T) {
	tgClient, _ := NewTgClient(&MockBotModel{}, &MockPriceModel{}, &MockScrapperModel{}, &config.Config{})

	msg1 := tgbotapi.Message{
		Text: "/help",
		From: &tgbotapi.User{UserName: "test"},
	}

	update := tgbotapi.Update{
		Message: &msg1,
	}

	tgClient.NewMessageReceived(update)
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

	assert.Equal(t, msg, expectedMsg)
}
