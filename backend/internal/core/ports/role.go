package ports

import model "game-server/internal/core/domain"

type Role interface {
	HasRole(roleType model.RoleType, viewer model.User) error
}
