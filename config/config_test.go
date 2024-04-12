package config

import (
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	// Mock environment variables
	os.Setenv("DB_PORT", "5432")
	os.Setenv("TG_DEBUG", "true")
	os.Setenv("PRICE_EXPIRATION_MIN", "5")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "dbname")
	os.Setenv("TG_TOKEN", "token")

	// Call the function
	config, err := ReadConfig()

	// Check for errors
	if err != nil {
		t.Errorf("ReadConfig() error = %v", err)
		return
	}

	// Check if the function correctly read the environment variables
	if config.Db.Host != "localhost" {
		t.Errorf("ReadConfig() = %v, want %v", config.Db.Host, "localhost")
	}
	if config.Db.User != "user" {
		t.Errorf("ReadConfig() = %v, want %v", config.Db.User, "user")
	}
	if config.Db.Password != "password" {
		t.Errorf("ReadConfig() = %v, want %v", config.Db.Password, "password")
	}
	if config.Db.DbName != "dbname" {
		t.Errorf("ReadConfig() = %v, want %v", config.Db.DbName, "dbname")
	}
	if config.Db.Port != 5432 {
		t.Errorf("ReadConfig() = %v, want %v", config.Db.Port, 5432)
	}
	if config.Tg.Token != "token" {
		t.Errorf("ReadConfig() = %v, want %v", config.Tg.Token, "token")
	}
	if config.Tg.Debug != true {
		t.Errorf("ReadConfig() = %v, want %v", config.Tg.Debug, true)
	}
	if config.App.ExpirationMin != 5 {
		t.Errorf("ReadConfig() = %v, want %v", config.App.ExpirationMin, 5)
	}
}
