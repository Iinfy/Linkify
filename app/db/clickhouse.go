package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var ch_conn driver.Conn

func ConnectClickhouse(url string) {
	opts, err := clickhouse.ParseDSN(url)
	if err != nil {
		log.Fatal(err)
	}

	tconn, err := clickhouse.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	if err = tconn.Ping(context.Background()); err != nil {
		log.Fatal("clickhouse ping: ", err)
	}
	ch_conn = tconn
}

func PrepareClickhouse() {
	err := ch_conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS stats(
		hash String,
	 	clicked_at DateTime64(3),
		user_agent String
		) ENGINE = MergeTree()
		ORDER BY clicked_at`)
	if err != nil {
		log.Fatal(err)
	}

}

func RecordClick(hash string, userAgent string) {
	batch, err := ch_conn.PrepareBatch(context.Background(), "INSERT INTO stats(hash, clicked_at, user_agent)")
	if err != nil {
		fmt.Println("prepare batch: ", err)
	}
	if err = batch.Append(hash, time.Now(), userAgent); err != nil {
		fmt.Println("batch append: ", err)
	}

	batch.Send()
}

func GetClicksByHash(hash string) uint64 {
	row := ch_conn.QueryRow(context.Background(),
		`SELECT
		count() as click_count
		FROM stats
		WHERE hash = ?`, hash)
	var click_count uint64
	if err := row.Scan(&click_count); err != nil {
		fmt.Println("row scan error")
		return 0
	}
	return click_count
}
