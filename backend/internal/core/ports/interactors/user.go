package interactors

import (
	"game-server/internal/core/domain"
	"time"
)

type UserInteractor interface {
	SignUp(user domain.User) error
	Get(user_id string) (*domain.User, error)
	Delete(user_id string) error
	Update(user_id string) error
	Block(user_id string, duration time.Time) error
}
