package pin

import (
	adapter "game-server/internal/adapter/repository"
	repo "game-server/internal/core/ports/repository"
)

type Interactor struct {
	Pin repo.PinRepository
}

func NewInteractor() *Interactor {
	return &Interactor{
		Pin: adapter.NerPinAdapter(),
	}
}
