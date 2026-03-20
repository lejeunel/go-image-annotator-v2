package assign_label

type FakePresenter struct {
	Got                        Response
	GotSuccess                 bool
	GotNotFoundErr             bool
	GotInternalErr             bool
	GotImageNotInCollectionErr bool
}

func (p *FakePresenter) Success(r Response) {
	p.Got = r
	p.GotSuccess = true
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrImageNotInCollection(error) {
	p.GotImageNotInCollectionErr = true
}
