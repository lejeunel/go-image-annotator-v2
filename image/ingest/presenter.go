package ingest

type FakePresenter struct {
	GotSuccess               bool
	GotCollectionNotFoundErr bool
	GotLabelNotFoundErr      bool
	GotInvalidImageDataErr   bool
	GotInternalErr           bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
}
func (p *FakePresenter) ErrCollectionNotFound(error) {
	p.GotCollectionNotFoundErr = true
}

func (p *FakePresenter) ErrInvalidImageData(error) {
	p.GotInvalidImageDataErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrLabelNotFound(error) {
	p.GotLabelNotFoundErr = true
}
