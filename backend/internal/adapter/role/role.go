package adapter

import (
	"game-server/internal/core/domain"
	"game-server/internal/core/ports"
)

type RolePort struct {
	db ports.DB
}

func New(db ports.DB) RolePort {
	return RolePort{db}
}

func (rp *RolePort) HasRole(roleType *domain.RoleType, viewer domain.User) error {
	roles, err := rp.db.Find()

}
