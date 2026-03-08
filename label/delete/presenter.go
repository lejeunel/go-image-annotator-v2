package delete

type DeleteOutputPort interface {
	ErrDependency(error)
	ErrInternal(error)
	Success()
}

type FakeDeletePresenter struct {
	GotDependencyErr bool
	GotInternalErr   bool
	GotSuccess       bool
}

func (p *FakeDeletePresenter) ErrDependency(error) {
	p.GotDependencyErr = true
}

func (p *FakeDeletePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakeDeletePresenter) Success() {
	p.GotSuccess = true
}
