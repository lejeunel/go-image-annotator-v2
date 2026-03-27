package unassign_label

import (
	"testing"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestHandleNotFoundErr(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, presenter, &im.FakeImageStore{Err: e.ErrNotFound})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected not found error")
	}
}

func TestHandleInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, presenter, &im.FakeImageStore{Err: e.ErrInternal})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestUnassignLabelNotAssignedToImageShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	image := im.NewImage(im.NewImageId(), *clc.NewCollection(clc.NewCollectionId(), "a-collection"))
	itr := NewInteractor(repo, presenter, &im.FakeImageStore{Return: image})
	itr.Execute(Request{Label: "label-not-assigned-to-image"})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected not found error")
	}
}

func TestHandleInternalErrOnRemoveLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	image := im.NewImage(im.NewImageId(), *clc.NewCollection(clc.NewCollectionId(), "a-collection"))
	image.AddLabel(lbl.NewLabel(lbl.NewLabelID(), "a-label"))
	repo := &FakeRepo{ErrOnRemoveLabel: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, presenter, &im.FakeImageStore{Return: image})
	itr.Execute(Request{Label: "a-label"})
	if presenter.GotSuccess || !presenter.GotInternalErr {
		t.Fatal("expected internal error")
	}
}
