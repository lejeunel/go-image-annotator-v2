package read_meta

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type FakePresenter struct {
	Got            im.ImageResponse
	GotInternalErr bool
	GotNotFoundErr bool
	GotSuccess     bool
}

func (p *FakePresenter) Success(r im.ImageResponse) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}
