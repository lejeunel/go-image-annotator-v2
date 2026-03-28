package update

type OutputPort interface {
	Success(Response)
	ErrDuplication(error)
	ErrNotFound(error)
	ErrInternal(error)
}

type FakePresenter struct {
	Got               Response
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

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
