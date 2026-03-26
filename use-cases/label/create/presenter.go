package create

type FakePresenter struct {
	Got               Response
	GotDuplicationErr bool
	GotInternalErr    bool
	GotSuccess        bool
	GotValidationErr  bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
func (p *FakePresenter) ErrDuplication(error) {
	p.GotDuplicationErr = true
}
func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}
