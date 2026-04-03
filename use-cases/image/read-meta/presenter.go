package read_meta

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type FakePresenter struct {
	Got            im.Response
	GotInternalErr bool
	GotNotFoundErr bool
	GotSuccess     bool
}

func (p *FakePresenter) Success(r im.Response) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}
