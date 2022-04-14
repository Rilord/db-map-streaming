package domain

import (
	"database/sql"
	"time"
)

type User struct {
	ID           string `json:"id"`
	SignedUpAt   time.Time
	UpdatedAt    sql.NullTime
	RemovedAt    sql.NullTime
	BlockedUntil sql.NullTime
}

type UserFilter struct {
	ID []string `json:"id"`
}

type UserConnection struct {
	PageInfo *PageInfo
}
