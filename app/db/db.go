package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
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
		log.Fatal(err)
	}
}

func AddUrl(url string, hash string) {
	_, err := conn.Exec(context.Background(), `INSERT INTO urls VALUES ($1,$2)`, hash, url)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUrlByHash(hash string) string {
	var url string
	err := conn.QueryRow(context.Background(), "SELECT url FROM urls WHERE hash = $1", hash).Scan(&url)
	if err != nil {
		log.Fatal()
		return ""
	}
	return url
}
