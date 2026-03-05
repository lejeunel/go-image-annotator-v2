package read

type ReadCollectionPresenter interface {
	Success(ReadCollectionResponse)
	ErrNotFound(string)
}

type FakeReadCollectionPresenter struct {
	Got            ReadCollectionResponse
	GotNotFoundErr bool
}

func (p *FakeReadCollectionPresenter) Success(r ReadCollectionResponse) {
	p.Got = r
}

func (p *FakeReadCollectionPresenter) ErrNotFound(m string) {
	p.GotNotFoundErr = true
}
