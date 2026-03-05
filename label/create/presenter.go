package create

type CreateLabelPresenter interface {
	Success(CreateLabelResponse)
	ErrDuplication(string)
	ErrInternal(string)
}

type FakeCreatePresenter struct {
	Got               CreateLabelResponse
	GotDuplicationErr bool
	GotInternalErr    bool
}

func (p *FakeCreatePresenter) Success(r CreateLabelResponse) {
	p.Got = r
}
func (p *FakeCreatePresenter) ErrDuplication(m string) {
	p.GotDuplicationErr = true
}
func (p *FakeCreatePresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}
