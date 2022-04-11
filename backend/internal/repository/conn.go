package repository

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() Database {
	pg, ok := os.LookupEnv("PGSQL")

	if !ok {
		log.Fatalln("Couldn't find PGSQL in local env")
	}

	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{})

	if err != nil {
		log.Fatalln("Couldn't open PgSQL db!")
	}

	return New(db)
}
