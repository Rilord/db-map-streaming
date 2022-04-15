package interactors

import "game-server/internal/core/domain"

type PinInteractor interface {
	Add(pin domain.Pin) error
	Get(pin_id string) (*domain.Pin, error)
	Delete(pin_id string) error
}
