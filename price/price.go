package price

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hemreari/goldwatcher/config"
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

type DbClient struct {
	Db *pgx.Conn
}

func NewDbClient(ctx context.Context, conf *config.Config) (*DbClient, error) {
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

func getDbConnUrl(conf *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.DbName,
	)
}

type PriceModel interface {
	InsertNewPrice(ctx context.Context, price *Price)
	GetLatestPrice(ctx context.Context, expirationMin int) *Price
}

// getInsertNewPriceSQL returns SQL that inserts to the prices table.
func getInsertNewPriceSQL() string {
	return fmt.Sprintf("INSERT INTO prices (%s, %s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		LastAtColName, Ayar22AltinColName, CeyrekColName, YarimColName, TamColName, CumhuriyetColName, IabKapanisColName)
}

// InsertPrice inserts given price information to prices table.
func (d *DbClient) InsertNewPrice(ctx context.Context, price *Price) {
	query := getInsertNewPriceSQL()

	cmdTag, err := d.Db.Exec(ctx,
		query, time.Now(), price.Ayar22Altin, price.Ceyrek, price.Yarim, price.Tam, price.Cumhuriyet, price.IabKapanis)
	if err != nil {
		log.Errorf("error while inserting new price record: %v", err)
	}

	log.Printf("cmd status: %v", cmdTag)
}

// getLatestPriceSQL returns SQL that selects from the price table where last_at
// bigger than the arg and orders by last in descending order.
func getLatestPriceSQL() string {
	return "SELECT * FROM prices WHERE last_at > $1 ORDER BY last_at DESC"
}

// getLatestPrice returns most recent price record according to threshold given
// in expirationMin param. If there is not any record in given threshold range
// returns nil.
// expirationMin is set in the config file.
func (d *DbClient) GetLatestPrice(ctx context.Context, expirationMin int) *Price {
	now := time.Now()
	then := now.Add(time.Duration(-expirationMin) * time.Minute)
	query := getLatestPriceSQL()

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
