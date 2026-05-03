package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var conn *pgx.Conn

func ConnectDatabase(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cconn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatal(err)
	}
	conn = cconn
	fmt.Println("db connected")
}

func PrepareDatabase() {
	_, err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS urls(
		hash text PRIMARY KEY,
		url TEXT NOT NULL UNIQUE)`)
	if err != nil {
		fmt.Println(err)
	}
	_, err = conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS stats(
		hash TEXT,
	 	clicked_at TIMESTAMP(3),
		user_agent TEXT)`)
	if err != nil {
		fmt.Println(err)
	}
}

func AddUrl(url string, hash string) {
	_, err := conn.Exec(context.Background(), `INSERT INTO urls VALUES ($1,$2)`, hash, url)
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
	err := conn.QueryRow(context.Background(), "SELECT url FROM urls WHERE hash = $1", hash).Scan(&url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return url
}

func RecordClick(hash string, userAgent string) {
	_, err := conn.Exec(context.Background(), `INSERT INTO stats(hash,clicked_at,user_agent)
		VALUES ($1,$2,$3)`, hash, time.Now(), userAgent)
	if err != nil {
		fmt.Println(err)
	}
}

func GetClicksByHash(hash string) int {
	var clicks int
	err := conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM stats WHERE hash=$1", hash).Scan(&clicks)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return clicks
}
