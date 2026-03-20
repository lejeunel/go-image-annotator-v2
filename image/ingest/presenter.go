package ingest

type FakePresenter struct {
	Got                      *Response
	GotSuccess               bool
	GotCollectionNotFoundErr bool
	GotLabelNotFoundErr      bool
	GotInvalidImageDataErr   bool
	GotDuplicateImage        bool
	GotInternalErr           bool
	GotValidationErr         bool
}

func (p *FakePresenter) Success(r Response) {
	p.Got = &r
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

func (p *FakePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}

func (p *FakePresenter) ErrDuplicateImage(error) {
	p.GotDuplicateImage = true
}
