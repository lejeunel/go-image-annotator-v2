package update

type UpdateCollectionPresenter interface {
	Success(UpdateCollectionResponse)
	ErrDuplication(string)
}

type FakeUpdateCollectionPresenter struct {
	Got               UpdateCollectionResponse
	GotDuplicationErr bool
}

func (p *FakeUpdateCollectionPresenter) ErrDuplication(m string) {
	p.GotDuplicationErr = true
}

func (p *FakeUpdateCollectionPresenter) Success(r UpdateCollectionResponse) {
	p.Got = r
}
