package annotator

import (
	"fmt"

	sto "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	"github.com/lejeunel/go-image-annotator-v2/application/scroller"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type AnnotatorState struct {
	Image    im.Image
	Scroller scroller.ScrollerState
}

type Annotator struct {
	imageId    im.ImageId
	collection string
	scroller   *scroller.Scroller
	store      *sto.ImageStore
}

func (a *Annotator) State() (*AnnotatorState, error) {
	image, err := a.store.Find(im.BaseImage{ImageId: a.imageId, Collection: a.collection})
	if err != nil {
		return nil, err
	}
	scrollerState, err := a.scroller.State()
	if err != nil {
		return nil, err
	}
	return &AnnotatorState{Image: *image, Scroller: *scrollerState}, nil
}

func NewAnnotator(scrollerRepo scroller.Repo, store *sto.ImageStore, imageId im.ImageId, collection string) (*Annotator, error) {
	scroller, err := scroller.New(scrollerRepo, imageId, scroller.WithCollection(collection))
	if err != nil {
		return nil, fmt.Errorf("building annotator: %w", err)
	}
	return &Annotator{scroller: scroller, store: store, imageId: imageId, collection: collection}, nil
}

type AnnotatorBuilder struct {
	repo  scroller.Repo
	store *sto.ImageStore
}

func (b *AnnotatorBuilder) Build(imageId im.ImageId, collection string) (*Annotator, error) {
	return NewAnnotator(b.repo, b.store, imageId, collection)
}

func NewAnnotatorBuilder(scrollerRepo scroller.Repo, store *sto.ImageStore) *AnnotatorBuilder {
	return &AnnotatorBuilder{repo: scrollerRepo, store: store}
}
