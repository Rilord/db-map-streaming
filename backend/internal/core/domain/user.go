package model

import (
	"database/sql"
)

type UserRole string

type User struct {
	ID           string `json:"id"`
	SignedUpAt   sql.NullTime
	updatedAt    sql.NullTime
	name         string
	email        string
	RemovedAt    sql.NullTime
	BlockedUntil sql.NullTime
}
