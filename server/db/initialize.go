package db

import (
	"game-server/log"
	"os"
	"time"

	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	dsn, ok := os.LookupEnv("POSTGRESQL_CONN")

	if !ok {
		log.Fatal("POSTGRESQL_CONN doesn't exist")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
