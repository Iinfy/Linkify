package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitPool(connString string) error {
	var err error
	pool, err = pgxpool.New(context.Background(), connString)
	return err
}

func PrepareDatabase() {
	_, err := pool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS urls(
		hash text PRIMARY KEY,
		url TEXT NOT NULL UNIQUE)`)
	if err != nil {
		fmt.Println(err)
	}
	_, err = pool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS stats(
		hash TEXT,
	 	clicked_at TIMESTAMP(3),
		user_agent TEXT)`)
	if err != nil {
		fmt.Println(err)
	}
}

func AddUrl(url string, hash string) {
	_, err := pool.Exec(context.Background(), `INSERT INTO urls VALUES ($1,$2)`, hash, url)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				fmt.Println("unique violation")
			}
		}

	}
}

func GetUrlByHash(hash string) string {
	var url string
	err := pool.QueryRow(context.Background(), "SELECT url FROM urls WHERE hash = $1", hash).Scan(&url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return url
}

func RecordClick(hash string, userAgent string) {
	_, err := pool.Exec(context.Background(), `INSERT INTO stats(hash,clicked_at,user_agent)
		VALUES ($1,$2,$3)`, hash, time.Now(), userAgent)
	if err != nil {
		fmt.Println(err)
	}
}

func GetClicksByHash(hash string) int {
	var clicks int
	err := pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM stats WHERE hash=$1", hash).Scan(&clicks)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return clicks
}
