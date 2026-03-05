package delete

type DeleteLabelPresenter interface {
	ErrDependency(string)
	ErrInternal(string)
	Success()
}

type FakeDeleteLabelPresenter struct {
	GotDependencyErr bool
	GotInternalErr   bool
	GotSuccess       bool
}

func (p *FakeDeleteLabelPresenter) ErrDependency(m string) {
	p.GotDependencyErr = true
}

func (p *FakeDeleteLabelPresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}

func (p *FakeDeleteLabelPresenter) Success() {
	p.GotSuccess = true
}
