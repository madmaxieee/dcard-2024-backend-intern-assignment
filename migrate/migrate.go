package main

import (
	"github.com/joho/godotenv"
	"github.com/madmaxieee/dcard-2024-backend-intern-assignment/database"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectToDb()
}

func main() {
	database.DB.AutoMigrate(&database.Advertisement{})
}
