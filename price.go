package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Price struct {
	Id           uint      `gorm:"unique;primaryKey;autoIncrement"`
	Last_at      time.Time `gorm:"autoCreateTime"`
	Ayar22_altin int
	Ceyrek       int
	Yarim        int
	Tam          int
	Cumhuriyet   int
	Iab_kapanis  int
}

func (d *DbClient) InsertNewPrice(price *Price) {
	result := d.Db.Create(price)
	err := result.Error
	if err != nil {
		log.Errorf("error while inserting new price record: %v", err)
	}

	rowsAffected := result.RowsAffected
	log.Print(rowsAffected)
}

func (d *DbClient) GetLatestPrice() *Price {
	return nil
}
