package create

type OutputPort interface {
	Success(Response)
	ErrDuplication(error)
	ErrInternal(error)
	ErrValidation(error)
}

type FakePresenter struct {
	Got               Response
	GotSuccess        bool
	GotDuplicationErr bool
	GotInternalErr    bool
	GotValidationErr  bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
func (p *FakePresenter) ErrDuplication(error) {
	p.GotDuplicationErr = true
}
func (p *FakePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
