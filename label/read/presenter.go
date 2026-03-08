package read

type ReadOutputPort interface {
	Success(ReadResponse)
	ErrNotFound(error)
	ErrInternal(error)
}

type FakeReadPresenter struct {
	Got            ReadResponse
	GotNotFoundErr bool
	GotInternalErr bool
	GotSuccess     bool
}

func (p *FakeReadPresenter) Success(r ReadResponse) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakeReadPresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakeReadPresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
