package adapter

import (
	"game-server/internal/core/ports"
)

type RolePort struct {
	db ports.DB
}

func New(db ports.DB) RolePort {
	return RolePort{db}
}
