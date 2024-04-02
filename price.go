package main

import (
	"context"
	"fmt"
	"time"

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

func (d *DbClient) InsertNewPrice(price *Price) {
	query := fmt.Sprintf("INSERT INTO prices (%s, %s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		LastAtColName, Ayar22AltinColName, CeyrekColName, YarimColName, TamColName, CumhuriyetColName, IabKapanisColName)

	cmdTag, err := d.Db.Exec(context.Background(),
		query, time.Now(), price.Ayar22Altin, price.Ceyrek, price.Yarim, price.Tam, price.Cumhuriyet, price.IabKapanis)
	if err != nil {
		log.Errorf("error while inserting new price record: %v", err)
	}

	log.Printf("cmd status: %v", cmdTag)
}

func (d *DbClient) GetLatestPrice() *Price {
	return nil
}
