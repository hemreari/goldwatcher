package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type DbClient struct {
	Db *pgx.Conn
}

func NewDbClient(ctx context.Context, conf *Config) *DbClient {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbConnUrl := getDbConnUrl(conf)
	conn, err := pgx.Connect(ctx, dbConnUrl)
	if err != nil {
		log.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &DbClient{Db: conn}
}

func getDbConnUrl(conf *Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.DbName,
	)
}
