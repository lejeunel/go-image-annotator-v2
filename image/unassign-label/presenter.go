package unassign_label

type FakePresenter struct {
	Got            Response
	GotSuccess     bool
	GotNotFoundErr bool
	GotInternalErr bool
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
