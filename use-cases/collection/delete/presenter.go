package delete

type OutputPort interface {
	ErrDependency(error)
	ErrInternal(error)
	ErrNotFound(error)
	Success()
}

type FakePresenter struct {
	GotDependencyErr bool
	GotInternalErr   bool
	GotNotFoundErr   bool
	GotSuccess       bool
}

func (p *FakePresenter) ErrDependency(error) {
	p.GotDependencyErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) Success() {
	p.GotSuccess = true
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}
