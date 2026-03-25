package read_raw

type FakePresenter struct {
	Got            Response
	GotNotFoundErr bool
	GotSuccess     bool
	GotInternalErr bool
}

func (p *FakePresenter) ErrNotFound(err error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(err error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
