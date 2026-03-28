package list

type OutputPort interface {
	Success(ListResponse)
	ErrInternal(error)
}

type FakePresenter struct {
	Got            ListResponse
	GotInternalErr bool
	GotSuccess     bool
}

func (p *FakePresenter) Success(r ListResponse) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
