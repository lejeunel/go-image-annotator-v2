package read

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrInternal(error)
}

type FakeReadPresenter struct {
	Got            Response
	GotNotFoundErr bool
	GotInternalErr bool
	GotSuccess     bool
}

func (p *FakeReadPresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakeReadPresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakeReadPresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
