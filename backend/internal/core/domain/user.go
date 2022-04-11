package domain

import (
	"database/sql"
	"time"
)

type UserRole string

type User struct {
	ID           string `json:"id"`
	SignedUpAt   time.Time
	UpdatedAt    sql.NullTime
	name         string
	email        string
	RemovedAt    sql.NullTime
	BlockedUntil sql.NullTime
}
