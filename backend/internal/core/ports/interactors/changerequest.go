package interactors

import "game-server/internal/core/domain"

type ChangeRequestInteractor interface {
	Add(request domain.ChangeRequest) error
	Get(request domain.ChangeRequest) (domain.ChangeRequest, error)
	Delete(request_id string) error
	Merge(request_id string) error
}
