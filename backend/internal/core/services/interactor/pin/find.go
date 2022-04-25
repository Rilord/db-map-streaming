package pin

import (
	"game-server/internal/core/domain"
	repo "game-server/internal/core/ports/repository"
)

func (i *Interactor) Find(filter repo.PinFilter) ([]*domain.Pin, error) {
	return i.Pin.Find(filter)
}
