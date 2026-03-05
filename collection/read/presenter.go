package read

type ReadCollectionPresenter interface {
	Success(ReadCollectionResponse)
	ErrNotFound(string)
	ErrInternal(string)
}

type FakeReadCollectionPresenter struct {
	Got            ReadCollectionResponse
	GotNotFoundErr bool
	GotInternalErr bool
}

func (p *FakeReadCollectionPresenter) Success(r ReadCollectionResponse) {
	p.Got = r
}

func (p *FakeReadCollectionPresenter) ErrNotFound(m string) {
	p.GotNotFoundErr = true
}

func (p *FakeReadCollectionPresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}
