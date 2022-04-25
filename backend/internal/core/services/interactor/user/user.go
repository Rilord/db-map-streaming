package user

import rp "game-server/internal/core/ports/repository"
import adapter "game-server/internal/adapter/repository"

type Interactor struct {
	User rp.UserRepository
}

func NewInteractor() *Interactor {
	return &Interactor{
		User: adapter.NewUserAdapter(),
	}
}
