package add_bbox

type FakePresenter struct {
	GotNotFoundErr   bool
	GotInternalErr   bool
	GotValidationErr bool
	GotSuccess       bool
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}

func (p *FakePresenter) Success(Response) {
	p.GotSuccess = true
}
