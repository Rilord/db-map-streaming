package repository

import (
	"database/sql"
	"fmt"
	"game-server/internal/core/domain"
	"game-server/internal/core/ports"
	"gorm.io/gorm"
	"time"
)

type userDB struct{ *Database }

func (d *Database) User() ports.UserDB { return &userDB{d} }

type User struct {
	ID           string    `gorm:"default:generated();"`
	SignedUpAt   time.Time `gorm:"default:now();"`
	UpdatedAt    sql.NullTime
	RemovedAt    sql.NullTime
	BlockedUntil sql.NullTime
}

func (usr *User) into() domain.User {
	return domain.User{
		ID:           usr.ID,
		SignedUpAt:   usr.SignedUpAt,
		UpdatedAt:    usr.UpdatedAt,
		RemovedAt:    usr.RemovedAt,
		BlockedUntil: usr.BlockedUntil,
	}

}

func (dst *User) copy(src *domain.User) {
	if src == nil {
		return
	}

	dst.ID = src.ID
	dst.SignedUpAt = src.SignedUpAt
	dst.UpdatedAt = src.UpdatedAt
	dst.RemovedAt = src.RemovedAt
	dst.BlockedUntil = src.BlockedUntil
}

func (db *userDB) Get(id string) (domain.User, error) {
	obj := User{}

	if err := db.db.Take(&obj, "id = ?", id).Error; err != nil {
		return domain.User{}, fmt.Errorf("take: %w", convertError(err))
	}

	return obj.into(), nil
}

func (db *userDB) Create(user *domain.User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}

	res := User{}
	res.copy(user)

	if err := db.db.Create(&res).Error; err != nil {
		return fmt.Errorf("create failed with error: %w", err)
	}

	*user = res.into()
	return nil
}
