package delete

import (
	"testing"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingResourceShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{Err: e.ErrNotFound}, &FakeRepo{})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected to get not found error")
	}

}

func TestHandleInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{Err: e.ErrInternal}, &FakeRepo{})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestDeleteNonExistingLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	id := im.NewImageId()
	collectionName := "a-collection"
	image := im.NewImage(id, *clc.NewCollection(clc.NewCollectionId(), collectionName))
	image.AddLabel(lbl.NewLabel(lbl.NewLabelID(), "a-label"))
	itr := NewInteractor(presenter, &im.FakeImageStore{Return: image}, &FakeRepo{ErrOnRemoveAnnotation: true, Err: e.ErrNotFound})
	itr.Execute(Request{})
	if presenter.GotSuccess || !(presenter.GotNotFoundErr) {
		t.Fatalf("expected not found error")
	}
}

func TestHandleInternalErrOnRemoveLabel(t *testing.T) {
	presenter := &FakePresenter{}
	id := im.NewImageId()
	collectionName := "a-collection"
	image := im.NewImage(id, *clc.NewCollection(clc.NewCollectionId(), collectionName))
	image.AddLabel(lbl.NewLabel(lbl.NewLabelID(), "a-label"))
	itr := NewInteractor(presenter, &im.FakeImageStore{Return: image}, &FakeRepo{ErrOnRemoveAnnotation: true, Err: e.ErrInternal})
	itr.Execute(Request{})
	if presenter.GotSuccess || !(presenter.GotInternalErr) {
		t.Fatalf("expected internal error")
	}
}

func TestDeleteNonExistingBoxShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	id := im.NewImageId()
	collectionName := "a-collection"
	image := im.NewImage(id, *clc.NewCollection(clc.NewCollectionId(), collectionName))
	box := a.NewBoundingBox(a.NewAnnotationId(), 1, 1, 1, 1,
		*lbl.NewLabel(lbl.NewLabelID(), "a-label"))
	image.AddBoundingBox(*box)
	itr := NewInteractor(presenter, &im.FakeImageStore{Return: image}, &FakeRepo{ErrOnRemoveAnnotation: true, Err: e.ErrNotFound})
	itr.Execute(Request{})
	if presenter.GotSuccess || !(presenter.GotNotFoundErr) {
		t.Fatalf("expected not found error")
	}
}

func TestHandleInternalErrOnDeleteBoxes(t *testing.T) {
	presenter := &FakePresenter{}
	id := im.NewImageId()
	collectionName := "a-collection"
	image := im.NewImage(id, *clc.NewCollection(clc.NewCollectionId(), collectionName))
	box := a.NewBoundingBox(a.NewAnnotationId(), 1, 1, 1, 1, *lbl.NewLabel(lbl.NewLabelID(), "a-label"))
	image.AddBoundingBox(*box)
	itr := NewInteractor(presenter, &im.FakeImageStore{Return: image}, &FakeRepo{ErrOnRemoveAnnotation: true, Err: e.ErrInternal})
	itr.Execute(Request{})
	if presenter.GotSuccess || !(presenter.GotInternalErr) {
		t.Fatalf("expected internal error")
	}
}
