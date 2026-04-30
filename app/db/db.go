package db

import (
	"log"

	"github.com/jackc/pgx"
)

var conn *pgx.Conn

func ConnectDatabase(url string) {
	connConfig, err := pgx.ParseConnectionString(url)
	if err != nil {
		log.Fatal(err)
	}
	conn, err = pgx.Connect(connConfig)

}

func PrepareDatabase() {
	_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS urls(
		hash varchar(32) PRIMARY KEY,
		url TEXT NOT NULL UNIQUE)`)
	if err != nil {
		log.Fatal(err)
	}
}

func AddUrl(url string, hash string) {
	_, err := conn.Exec(`INSERT INTO urls VALUES ($1,$2)`, hash, url)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUrlByHash(hash string) string {
	var url string
	err := conn.QueryRow("SELECT url FROM urls WHERE hash = $1", hash).Scan(&url)
	if err != nil {
		log.Fatal()
		return ""
	}
	return url
}
