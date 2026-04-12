package annotator

import (
	"errors"
	"testing"

	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	an "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	addbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/add-bbox"
	updbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/modify-bbox"
	del "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/remove"
)

func TestHandleErrorOnScrollerInit(t *testing.T) {
	view := &FakeView{}
	a := &Annotator{scroller: &FakeScroller{ErrOnInit: true, Err: e.ErrInternal},
		store: &FakeStore{}}
	a.Init(im.NewImageId(), "a-collection", view)
	if !errors.Is(view.GotErr, e.ErrInternal) {
		t.Fatal("expected to handle error on scroller init")
	}
}

func TestInitializeScrollerOnInit(t *testing.T) {
	scroller := &FakeScroller{}
	a := &Annotator{scroller: scroller,
		store: &FakeStore{}}
	a.scroller = scroller
	a.Init(im.NewImageId(), "a-collection", &FakeView{})
	if !scroller.IsInit {
		t.Fatal("expected to initialize scroller")
	}
}

func TestDrawScrollerOnInit(t *testing.T) {
	view := &FakeView{}
	a := &Annotator{scroller: &FakeScroller{}, store: &FakeStore{}}
	a.Init(im.NewImageId(), "a-collection", view)
	if !view.DrewScroller {
		t.Fatal("expected to draw scroller")
	}
}
func TestHandleErrorOnRetrieveImage(t *testing.T) {
	view := &FakeView{}
	a := &Annotator{scroller: &FakeScroller{},
		store: &FakeStore{ErrOnFind: true, Err: e.ErrInternal}}
	a.Init(im.NewImageId(), "collection", view)
	if !errors.Is(view.GotErr, e.ErrInternal) {
		t.Fatalf("expected internal error got %v", view.GotErr)
	}
}

func createAnnotator() (*Annotator, *im.Image, *FakeView) {
	view := &FakeView{}
	image := im.NewImage(im.NewImageId(),
		*clc.NewCollection(clc.NewCollectionId(), "name"))
	label := lbl.NewLabel(lbl.NewLabelId(), "a-label")
	box := a.NewBoundingBox(a.NewAnnotationId(), 1, 1, 1, 1, *label)
	image.AddBoundingBox(*box)
	annotator := &Annotator{scroller: &FakeScroller{},
		store: &FakeStore{Return: *image}}
	return annotator, image, view

}

func TestDrawImageOnInit(t *testing.T) {
	annotator, image, view := createAnnotator()
	annotator.Init(image.Id, image.Collection.Name, view)
	if view.DrewImage.Id != image.Id {
		t.Fatalf("expected to draw image with id %v, got %v",
			image.Id, view.DrewImage.Id)
	}
}

func TestShowImageInfoOnInit(t *testing.T) {
	annotator, image, view := createAnnotator()
	annotator.Init(image.Id, image.Collection.Name, view)
	if view.DrewImageInfo.Id != image.Id {
		t.Fatalf("expected to show image info with id %v, got %v",
			image.Id, view.DrewImageInfo.Id)
	}
}

func TestShowAnnotationsOnInit(t *testing.T) {
	annotator, image, view := createAnnotator()
	annotator.Init(image.Id, image.Collection.Name, view)
	if view.DrewAnnotations[0].Id != image.BoundingBoxes[0].Id {
		t.Fatalf("expected to show annotation with id %v, got %v",
			image.BoundingBoxes[0].Id, view.DrewAnnotations[0].Id)
	}
}

func TestAddBox(t *testing.T) {
	view := &FakeView{}
	adder := &FakeBoxAdder{}
	annotator := &Annotator{boxAdder: adder}
	imageId := im.NewImageId()
	annotator.AddBox(addbox.Request{ImageId: imageId}, view)
	if adder.Got.ImageId != imageId {
		t.Fatalf("expected to add bbox on image %v, got %v",
			imageId, adder.Got.ImageId)
	}
}

func TestUpdateBox(t *testing.T) {
	view := &FakeView{}
	updater := &FakeBoxUpdater{}
	annotator := &Annotator{boxUpdater: updater}
	annotationId := an.NewAnnotationId()
	annotator.UpdateBox(updbox.Request{AnnotationId: annotationId}, view)
	if updater.Got.AnnotationId != annotationId {
		t.Fatalf("expected to modify bbox with id %v, got %v",
			annotationId, updater.Got.AnnotationId)
	}
}

func TestRemoveBox(t *testing.T) {
	view := &FakeView{}
	deleter := &FakeBoxDeleter{}
	annotator := &Annotator{boxDeleter: deleter}
	annotationId := an.NewAnnotationId()
	annotator.DeleteBox(del.Request{Id: annotationId}, view)
	if deleter.Got.Id != annotationId {
		t.Fatalf("expected to delete bbox with id %v, got %v",
			annotationId, deleter.Got.Id)
	}
}
