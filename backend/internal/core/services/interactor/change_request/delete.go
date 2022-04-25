package changerequest

func (i *Interactor) Delete(request_id string) error {
	return i.Request.Delete(request_id)
}
