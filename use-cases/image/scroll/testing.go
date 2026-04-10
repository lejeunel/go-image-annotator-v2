package scroll

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	t "github.com/lejeunel/go-image-annotator-v2/shared/testing"
)

type FakePresenter struct {
	Got        *im.Image
	GotSuccess bool
	t.TestingErrPresenter
}

func (p *FakePresenter) Success(r *im.Image) {
	p.GotSuccess = true
	p.Got = r
}

type FakeImageRepo struct {
	Err          error
	ReturnedNext bool
	ReturnedPrev bool
}

func (r *FakeImageRepo) GetAdjacent(id im.ImageId, collection string, d ScrollingDirection) (*im.BaseImage, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	if d == ScrollNext {
		r.ReturnedNext = true
	} else {
		r.ReturnedPrev = true
	}

	return &im.BaseImage{}, nil
}
