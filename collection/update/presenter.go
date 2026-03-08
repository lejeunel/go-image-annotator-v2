package update

type UpdateOutputPort interface {
	Success(UpdateResponse)
	ErrDuplication(error)
	ErrNotFound(error)
	ErrInternal(error)
}

type FakeUpdatePresenter struct {
	Got               UpdateResponse
	GotDuplicationErr bool
	GotNotFoundErr    bool
	GotInternalErr    bool
	GotSuccess        bool
}

func (p *FakeUpdatePresenter) ErrDuplication(error) {
	p.GotDuplicationErr = true
}

func (p *FakeUpdatePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakeUpdatePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakeUpdatePresenter) Success(r UpdateResponse) {
	p.GotSuccess = true
	p.Got = r
}
