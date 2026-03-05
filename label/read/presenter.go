package read

type ReadPresenter interface {
	Success(ReadResponse)
	ErrNotFound(string)
	ErrInternal(string)
}

type FakeReadPresenter struct {
	Got            ReadResponse
	GotNotFoundErr bool
	GotInternalErr bool
}

func (p *FakeReadPresenter) Success(r ReadResponse) {
	p.Got = r
}

func (p *FakeReadPresenter) ErrNotFound(m string) {
	p.GotNotFoundErr = true
}

func (p *FakeReadPresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}
