package main

import (
	"fmt"
	"linkify/db"
	"linkify/handlers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env load failed")
	}
	corsHost := os.Getenv("CORS_HOST")
	if corsHost == "" {
		panic("CORS_HOST is not set")
	}
	db.ConnectDatabase(os.Getenv("POSTGRES_URL"))
	db.PrepareDatabase()
	db.ConnectClickhouse(os.Getenv("CLICKHOUSE_URL"))
	db.PrepareClickhouse()
	handlers.StartServer(":8090")
}
