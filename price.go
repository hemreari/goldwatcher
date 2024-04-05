package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type PriceColumnName string

const (
	LastAtColName      PriceColumnName = "last_at"
	Ayar22AltinColName PriceColumnName = "ayar22_altin"
	CeyrekColName      PriceColumnName = "ceyrek"
	YarimColName       PriceColumnName = "yarim"
	TamColName         PriceColumnName = "tam"
	CumhuriyetColName  PriceColumnName = "cumhuriyet"
	IabKapanisColName  PriceColumnName = "iab_kapanis"
)

type Price struct {
	Id          uint
	LastAt      time.Time
	Ayar22Altin int
	Ceyrek      int
	Yarim       int
	Tam         int
	Cumhuriyet  int
	IabKapanis  int
}

// InsertNewPrice inserts given price information to prices table.
func (d *DbClient) InsertNewPrice(ctx context.Context, price *Price) {
	query := fmt.Sprintf("INSERT INTO prices (%s, %s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		LastAtColName, Ayar22AltinColName, CeyrekColName, YarimColName, TamColName, CumhuriyetColName, IabKapanisColName)

	cmdTag, err := d.Db.Exec(ctx,
		query, time.Now(), price.Ayar22Altin, price.Ceyrek, price.Yarim, price.Tam, price.Cumhuriyet, price.IabKapanis)
	if err != nil {
		log.Errorf("error while inserting new price record: %v", err)
	}

	log.Printf("cmd status: %v", cmdTag)
}

// GetLatestPrice returns most recent price record according to threshold given
// in expirationMin param. If there is not any record in given threshold range
// returns nil.
// expirationMin is set in the config file.
func (d *DbClient) GetLatestPrice(ctx context.Context, expirationMin int) *Price {
	now := time.Now()
	then := now.Add(time.Duration(-expirationMin) * time.Minute)
	log.Printf("then: %v", then)
	query := "SELECT * FROM prices WHERE last_at > $1 ORDER BY last_at DESC"

	var err error
	rows, err := d.Db.Query(ctx, query, then)
	if err != nil {
		log.Errorf("error while getting latest price: %v", err)
		return nil
	}
	price, err := pgx.CollectOneRow[Price](rows, pgx.RowToStructByName[Price])
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Warnf("couldn't find any record newer than %d mins.", expirationMin)
			return nil
		}
		log.Errorf("error while getting latest price: %v", err)
	}

	return &price
}
