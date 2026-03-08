package update

type UpdateOutputPort interface {
	Success(UpdateResponse)
	ErrDuplication(string)
	ErrNotFound(string)
	ErrInternal(string)
}

type FakeUpdatePresenter struct {
	Got               UpdateResponse
	GotDuplicationErr bool
	GotNotFoundErr    bool
	GotInternalErr    bool
	GotSuccess        bool
}

func (p *FakeUpdatePresenter) ErrDuplication(m string) {
	p.GotDuplicationErr = true
}

func (p *FakeUpdatePresenter) ErrNotFound(m string) {
	p.GotNotFoundErr = true
}

func (p *FakeUpdatePresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}

func (p *FakeUpdatePresenter) Success(r UpdateResponse) {
	p.GotSuccess = true
	p.Got = r
}
