package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
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
	cfg := readConfig()
	db := NewDbStruct(cfg)

	NewTgStruct(cfg)

	price := GetPrices()
	db.InsertNewPrice(&price)
}
