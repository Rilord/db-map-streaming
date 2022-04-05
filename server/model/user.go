package model

import (
	"database/sql"
	"time"
)

type UserRole string

const (
	UserRoleContributor UserRole = "CONTRIBUTOR"
	UserRoleModerator   UserRole = "MODERATOR"
	UserRoleViewer      UserRole = "VIEWER"
)

type User struct {
	ID           string    `gorm:default:generated();" json:"id"`
	SignedUpAt   time.Time `gorm:default:now();"`
	RemovedAt    sql.NullTime
	BlockedUntil sql.NullTime
	role         UserRole
}
