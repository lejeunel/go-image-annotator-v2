package update

type UpdateCollectionPresenter interface {
	Success(UpdateCollectionResponse)
	ErrDuplication(string)
	ErrNotFound(string)
	ErrInternal(string)
}

type FakeUpdateCollectionPresenter struct {
	Got               UpdateCollectionResponse
	GotDuplicationErr bool
	GotNotFoundErr    bool
	GotInternalErr    bool
}

func (p *FakeUpdateCollectionPresenter) ErrDuplication(m string) {
	p.GotDuplicationErr = true
}

func (p *FakeUpdateCollectionPresenter) ErrNotFound(m string) {
	p.GotNotFoundErr = true
}

func (p *FakeUpdateCollectionPresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}

func (p *FakeUpdateCollectionPresenter) Success(r UpdateCollectionResponse) {
	p.Got = r
}
