package pin

func (i *Interactor) Delete(pin_id string) error {
	return i.Pin.Delete(pin_id)
}
