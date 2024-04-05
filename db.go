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

func NewDbClient(ctx context.Context, conf *Config) (*DbClient, error) {
	dbConnUrl := getDbConnUrl(conf)
	conn, err := pgx.Connect(ctx, dbConnUrl)
	if err != nil {
		log.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't establish connection to db: %v", err)
	}

	log.Print("db connection is established.")

	return &DbClient{Db: conn}, nil
}

func getDbConnUrl(conf *Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.DbName,
	)
}
