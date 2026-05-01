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
	db.ConnectDatabase(os.Getenv("POSTGRES_URL"))
	db.PrepareDatabase()
	handlers.StartServer(":" + os.Getenv("PORT"))
}
