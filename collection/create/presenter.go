package create

type FakeCreateCollectionPresenter struct {
	Got               CreateCollectionResponse
	GotDuplicationErr bool
}

func (p *FakeCreateCollectionPresenter) Success(r CreateCollectionResponse) {
	p.Got = r
}
func (p *FakeCreateCollectionPresenter) ErrDuplication(m string) {
	p.GotDuplicationErr = true
}

type CreateCollectionPresenter interface {
	Success(CreateCollectionResponse)
	ErrDuplication(string)
}
