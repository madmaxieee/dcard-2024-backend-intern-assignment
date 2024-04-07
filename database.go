package main

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	DB, err = gorm.Open(sqlite.Dialector{DSN: os.Getenv("DATABASE_URL"), DriverName: "libsql"}, &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
		panic(err)
	}
}
