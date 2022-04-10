package db

import (
	"game-server/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Setup() *gorm.DB {
	dsn, ok := os.LookupEnv("POSTGRESQL_CONN")

	if !ok {
		log.Fatal("POSTGRESQL_CONN doesn't exist")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return database
}
