package read

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrInternal(error)
}

type FakePresenter struct {
	Got            Response
	GotNotFoundErr bool
	GotInternalErr bool
	GotSuccess     bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
