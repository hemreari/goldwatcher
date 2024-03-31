package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbClient struct {
	Db *gorm.DB
}

func NewDbClient(conf *Config) *DbClient {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  getDbDSNString(conf),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Price{})
	return &DbClient{Db: db}
}

func getDbDSNString(conf *Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.DbName,
	)
}
