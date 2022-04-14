package repository

import (
	"database/sql"
	"game-server/internal/core/domain"
	"time"
)

type Role struct {
	Type      domain.RoleType
	UserID    string
	CreatedAt time.Time
	RemovedAt sql.NullTime
}

func (r *Role) into()
