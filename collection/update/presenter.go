package update

type UpdateOutputPort interface {
	Success(UpdateResponse)
	ErrDuplication(error)
	ErrNotFound(error)
	ErrInternal(error)
}

type FakePresenter struct {
	Got               UpdateResponse
	GotDuplicationErr bool
	GotNotFoundErr    bool
	GotInternalErr    bool
	GotSuccess        bool
}

func (p *FakePresenter) ErrDuplication(error) {
	p.GotDuplicationErr = true
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) Success(r UpdateResponse) {
	p.GotSuccess = true
	p.Got = r
}
