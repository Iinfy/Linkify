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
	 	clicked_at DateTime64(3)
		) ENGINE = MergeTree()
		ORDER BY clicked_at`)
	if err != nil {
		log.Fatal(err)
	}

}

func RecordClick(hash string) {
	batch, err := ch_conn.PrepareBatch(context.Background(), "INSERT INTO stats(hash, clicked_at)")
	if err != nil {
		fmt.Println("prepare batch: ", err)
	}
	if err = batch.Append(hash, time.Now()); err != nil {
		fmt.Println("batch append: ", err)
	}

	batch.Send()
}
