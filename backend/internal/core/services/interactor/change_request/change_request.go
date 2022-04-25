package changerequest

import (
	adapter "game-server/internal/adapter/repository"
	repo "game-server/internal/core/ports/repository"
)

type Interactor struct {
	Request repo.ChangeRequestRepository
}

func NewInteractor() *Interactor {
	return &Interactor{
		Request: adapter.NewChangeRequestAdapter(),
	}

}
