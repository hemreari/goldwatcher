package price

import (
	"context"
	"fmt"
	"testing"

	"github.com/hemreari/goldwatcher/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDbClient struct {
	mock.Mock
}

func (m *MockDbClient) InsertNewPrice(ctx context.Context, price *Price) {
	m.Called(ctx, price)
}

func (m *MockDbClient) GetLatestPrice(ctx context.Context, expirationMin int) *Price {
	args := m.Called(ctx, expirationMin)
	return args.Get(0).(*Price)
}

func TestGetDbConnUrl(t *testing.T) {
	conf := &config.Config{
		Db: config.DbConf{
			User:     "testUser",
			Password: "testPassword",
			Host:     "localhost",
			Port:     5432,
			DbName:   "testDb",
		},
	}

	expected := "postgres://testUser:testPassword@localhost:5432/testDb"
	actual := getDbConnUrl(conf)

	assert.Equal(t, actual, expected)
}

func TestInsertNewPrice(t *testing.T) {
	mockDbClient := new(MockDbClient)
	price := &Price{Ayar22Altin: 100, Ceyrek: 200, Yarim: 300, Tam: 400, Cumhuriyet: 500, IabKapanis: 600}

	mockDbClient.On("InsertNewPrice", mock.Anything, price).Return()

	// Call the function with the mock
	mockDbClient.InsertNewPrice(context.Background(), price)

	// Assert that the expectations were met
	mockDbClient.AssertExpectations(t)
}

func TestGetLatestPrice(t *testing.T) {
	mockDbClient := new(MockDbClient)
	price := &Price{Ayar22Altin: 100, Ceyrek: 200, Yarim: 300, Tam: 400, Cumhuriyet: 500, IabKapanis: 600}

	mockDbClient.On("GetLatestPrice", mock.Anything, 30).Return(price)

	// Call the function with the mock
	result := mockDbClient.GetLatestPrice(context.Background(), 30)

	// Assert that the expectations were met
	mockDbClient.AssertExpectations(t)

	// Assert that the result is what we expect
	assert.Equal(t, result, price)
}

func TestGetInsertNewPriceSQL(t *testing.T) {
	expectedSQL := fmt.Sprintf("INSERT INTO prices (%s, %s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		LastAtColName, Ayar22AltinColName, CeyrekColName, YarimColName, TamColName, CumhuriyetColName, IabKapanisColName)

	gotSQL := getInsertNewPriceSQL()

	assert.Equal(t, gotSQL, expectedSQL)
}

func TestGetLatestPriceSQL(t *testing.T) {
	expected := "SELECT * FROM prices WHERE last_at > $1 ORDER BY last_at DESC"
	result := getLatestPriceSQL()

	assert.Equal(t, result, expected)
}
