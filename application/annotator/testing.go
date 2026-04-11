package annotator

import (
	scr "github.com/lejeunel/go-image-annotator-v2/application/scroller"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type FakeScroller struct {
	Err       error
	ErrOnInit bool
	IsInit    bool
}

func (s *FakeScroller) Init(imageId im.ImageId, opts ...scr.Option) (*scr.ScrollerState, error) {
	if s.ErrOnInit {
		return nil, s.Err
	}
	s.IsInit = true
	return &scr.ScrollerState{}, nil
}

type FakeView struct {
	DrewScroller bool
	DrewImage    *im.Image
	GotErr       error
}

func (v *FakeView) Error(err error) {
	v.GotErr = err
}

func (v *FakeView) DrawScroller(s scr.ScrollerState) {
	v.DrewScroller = true
}

func (v *FakeView) DrawImage(image im.Image) {
	v.DrewImage = &image
}

type FakeStore struct {
	Return    im.Image
	ErrOnFind bool
	Err       error
}

func (s *FakeStore) Find(im.BaseImage) (*im.Image, error) {
	if s.ErrOnFind {
		return nil, s.Err
	}
	return &s.Return, nil
}
