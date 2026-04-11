package annotator

import (
	sto "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	"github.com/lejeunel/go-image-annotator-v2/application/scroller"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type AnnotatorState struct {
	Image    im.Image
	Scroller scroller.ScrollerState
}

type AnnotatorView interface {
	DrawScroller(scroller.ScrollerState)
	DrawImage(im.Image)
	Error(error)
}

type Annotator struct {
	imageId    im.ImageId
	collection string
	Scroller   scroller.Interface
	Store      sto.Interface
}

func (a *Annotator) Init(imageId im.ImageId, collection string, view AnnotatorView) {
	scrollerState, err := a.Scroller.Init(imageId, scroller.WithCollection(collection))
	if err != nil {
		view.Error(err)
		return
	}
	image, err := a.Store.Find(im.BaseImage{ImageId: imageId, Collection: collection})
	if err != nil {
		view.Error(err)
		return
	}
	view.DrawScroller(*scrollerState)
	view.DrawImage(*image)
}

func NewAnnotator(scrollerRepo scroller.Repo, store *sto.ImageStore) *Annotator {
	return &Annotator{Scroller: scroller.New(scrollerRepo),
		Store: store}
}
