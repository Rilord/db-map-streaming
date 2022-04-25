package changerequest

import "game-server/internal/core/domain"

func (i *Interactor) Get(request_id string) (domain.ChangeRequest, error) {
	return i.Request.Get(request_id)
}
