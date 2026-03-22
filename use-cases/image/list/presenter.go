package list

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrInternal(error)
}

type FakePresenter struct {
	Got            Response
	GotInternalErr bool
	GotNotFoundErr bool
	GotSuccess     bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}
