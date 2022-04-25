package pin

import "game-server/internal/core/domain"

func (i *Interactor) Update(pin domain.Pin) error {
	return i.Pin.Update(pin)
}
