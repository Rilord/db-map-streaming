package pininteractor

import (
	"game-server/internal/adapter"
	repository "game-server/internal/core/ports/repository"
)

type Interactor struct {
	pin repository.PinRepository
}

func New() *Interactor {
	interactor := Interactor{
		pin: adapter.NewPinInteractorAdapter(),
	}

	return &interactor
}
