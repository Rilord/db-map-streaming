package model

import (
	"database/sql"
	"time"
)

type UserRole string

type User struct {
	ID           string    `gorm:default:generated();" json:"id"`
	SignedUpAt   time.Time `gorm:default:now();"`
	updatedAt    time.Time
	name         string
	email        string
	RemovedAt    sql.NullTime
	BlockedUntil sql.NullTime
}
