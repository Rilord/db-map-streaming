package pininteractor

import "game-server/internal/core/domain"

func (i *Interactor) Add(pin domain.Pin) error {
	return i.pin.Add(pin)
}
