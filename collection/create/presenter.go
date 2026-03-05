package create

type CreateCollectionPresenter interface {
	Success(CreateCollectionResponse)
	ErrDuplication(string)
	ErrInternal(string)
}

type FakeCreateCollectionPresenter struct {
	Got               CreateCollectionResponse
	GotDuplicationErr bool
	GotInternalErr    bool
}

func (p *FakeCreateCollectionPresenter) Success(r CreateCollectionResponse) {
	p.Got = r
}
func (p *FakeCreateCollectionPresenter) ErrDuplication(m string) {
	p.GotDuplicationErr = true
}
func (p *FakeCreateCollectionPresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}
