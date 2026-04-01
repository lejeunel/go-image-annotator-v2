package list

type OutputPort interface {
	Success(Response)
	ErrInternal(error)
}

type FakePresenter struct {
	Got            Response
	GotInternalErr bool
	GotSuccess     bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
