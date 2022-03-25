package db

import (
	"log"
	"os"
	"server/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn, ok := os.LookupEnv("POSTGRES_CONNECTION")

	if !ok {
		log.Fatal("envar POSTGRES_CONNECTION doesn't exist")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database")
	}

	DB = database

	return database

}
