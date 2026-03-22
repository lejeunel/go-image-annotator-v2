package unassign_label

type FakePresenter struct {
	GotSuccess       bool
	GotNotFoundErr   bool
	GotInternalErr   bool
	GotDependencyErr bool
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrDependency(error) {
	p.GotDependencyErr = true
}

func (p *FakePresenter) Success(Response) {
	p.GotSuccess = true
}
