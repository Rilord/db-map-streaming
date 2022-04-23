package pin

import "game-server/internal/core/domain"

func (i *Interactor) Add(pin domain.Pin) error {
	return i.Pin.Add(pin)
}
