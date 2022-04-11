package repository

import "os"

func Connect() Database {
	db, ok := os.LookupEnv("PGSQL")
}
