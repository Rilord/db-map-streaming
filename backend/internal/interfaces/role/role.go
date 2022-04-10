package role

import "game-server/model"

type Interface interface {
	HasRole(RoleType model.RoleType, viewer model.User)
}
