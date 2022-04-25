package changerequest

import "game-server/internal/core/domain"

func (i *Interactor) Add(request domain.ChangeRequest) error {
	return i.Request.Add(request)
}
