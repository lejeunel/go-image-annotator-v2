package list

type ListOutputPort interface {
	Success(ListResponse)
	ErrInternal(error)
}

type FakeListPresenter struct {
	Got            ListResponse
	GotInternalErr bool
	GotSuccess     bool
}

func (p *FakeListPresenter) Success(r ListResponse) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakeListPresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
