package read_meta

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	t "github.com/lejeunel/go-image-annotator-v2/shared/testing"
)

type FakePresenter struct {
	Got        im.Response
	GotSuccess bool
	t.TestingErrPresenter
}

func (p *FakePresenter) Success(r im.Response) {
	p.GotSuccess = true
	p.Got = r
}
