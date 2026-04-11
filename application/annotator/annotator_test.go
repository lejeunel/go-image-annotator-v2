package annotator

import (
	"errors"
	"testing"

	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func TestHandleErrorOnScrollerInit(t *testing.T) {
	view := &FakeView{}
	a := &Annotator{Scroller: &FakeScroller{ErrOnInit: true, Err: e.ErrInternal},
		Store: &FakeStore{}}
	a.Init(im.NewImageId(), "a-collection", view)
	if !errors.Is(view.GotErr, e.ErrInternal) {
		t.Fatal("expected to handle error on scroller init")
	}
}

func TestInitializeScrollerOnInit(t *testing.T) {
	scroller := &FakeScroller{}
	a := &Annotator{Scroller: scroller,
		Store: &FakeStore{}}
	a.Scroller = scroller
	a.Init(im.NewImageId(), "a-collection", &FakeView{})
	if !scroller.IsInit {
		t.Fatal("expected to initialize scroller")
	}
}

func TestDrawScrollerOnInit(t *testing.T) {
	view := &FakeView{}
	a := &Annotator{Scroller: &FakeScroller{}, Store: &FakeStore{}}
	a.Init(im.NewImageId(), "a-collection", view)
	if !view.DrewScroller {
		t.Fatal("expected to draw scroller")
	}
}
func TestHandleErrorOnRetrieveImage(t *testing.T) {
	view := &FakeView{}
	a := &Annotator{Scroller: &FakeScroller{},
		Store: &FakeStore{ErrOnFind: true, Err: e.ErrInternal}}
	a.Init(im.NewImageId(), "collection", view)
	if !errors.Is(view.GotErr, e.ErrInternal) {
		t.Fatalf("expected internal error got %v", view.GotErr)
	}
}

func TestDrawImageOnInit(t *testing.T) {
	view := &FakeView{}
	image := im.NewImage(im.NewImageId(),
		*clc.NewCollection(clc.NewCollectionId(), "name"))
	a := &Annotator{Scroller: &FakeScroller{},
		Store: &FakeStore{Return: *image}}
	a.Init(image.Id, image.Collection.Name, view)
	if view.DrewImage.Id != image.Id {
		t.Fatalf("expected to draw image with id %v, got %v",
			image.Id, view.DrewImage.Id)
	}
}
