package delete

type DeletePresenter interface {
	ErrDependency(string)
	ErrInternal(string)
	Success()
}

type FakeDeletePresenter struct {
	GotDependencyErr bool
	GotInternalErr   bool
	GotSuccess       bool
}

func (p *FakeDeletePresenter) ErrDependency(m string) {
	p.GotDependencyErr = true
}

func (p *FakeDeletePresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}

func (p *FakeDeletePresenter) Success() {
	p.GotSuccess = true
}
