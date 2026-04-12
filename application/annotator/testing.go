package annotator

import (
	scr "github.com/lejeunel/go-image-annotator-v2/application/scroller"
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	addbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/add-bbox"
	updbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/modify-bbox"
	del "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/remove"
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
	DrewScroller    bool
	DrewImage       *im.Image
	DrewImageInfo   *ImageInfo
	DrewAnnotations []a.Annotation
	GotErr          error
}

func (v *FakeView) Error(err error) {
	v.GotErr = err
}

func (v *FakeView) DrawScroller(s scr.ScrollerState) {
	v.DrewScroller = true
}
func (v *FakeView) DrawAnnotationList(annotations []a.Annotation) {
	v.DrewAnnotations = annotations
}

func (v *FakeView) DrawImage(image im.Image) {
	v.DrewImage = &image
}
func (v *FakeView) DrawImageInfo(info ImageInfo) {
	v.DrewImageInfo = &info
}

func (v *FakeView) SuccessAddBox(r addbox.Response) {
}
func (v *FakeView) SuccessUpdateBox(r updbox.Response) {
}
func (v *FakeView) SuccessDeleteAnnotation(r del.Response) {
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

type FakeBoxAdder struct {
	Got addbox.Request
}

func (b *FakeBoxAdder) Execute(r addbox.Request, o addbox.OutputPort) {
	b.Got = r
}

type FakeBoxUpdater struct {
	Got updbox.Request
}

func (b *FakeBoxUpdater) Execute(r updbox.Request, o updbox.OutputPort) {
	b.Got = r
}

type FakeBoxDeleter struct {
	Got del.Request
}

func (b *FakeBoxDeleter) Execute(r del.Request, o del.OutputPort) {
	b.Got = r
}
