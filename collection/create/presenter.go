package create

type CreateOutputPort interface {
	Success(CreateResponse)
	ErrDuplication(error)
	ErrInternal(error)
	ErrValidation(error)
}

type FakeCreatePresenter struct {
	Got               CreateResponse
	GotSuccess        bool
	GotDuplicationErr bool
	GotInternalErr    bool
	GotValidationErr  bool
}

func (p *FakeCreatePresenter) Success(r CreateResponse) {
	p.GotSuccess = true
	p.Got = r
}
func (p *FakeCreatePresenter) ErrDuplication(error) {
	p.GotDuplicationErr = true
}
func (p *FakeCreatePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}

func (p *FakeCreatePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
