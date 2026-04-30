package main

import (
	"PastebinBack/app/db"
	"PastebinBack/app/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectDatabase(os.Getenv("POSTGRES_URL"))
	db.PrepareDatabase()
	handlers.StartServer(":" + os.Getenv("PORT"))
}
