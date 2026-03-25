package read_meta

import (
	"testing"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestHandleNotFoundError(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{Err: e.ErrNotFound})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{Err: e.ErrInternal})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestFindImage(t *testing.T) {
	presenter := &FakePresenter{}
	existingImage := im.NewImage(im.NewImageId(), clc.NewCollection("a-collection"))
	existingImage.AddLabel(lbl.NewLabel("a-label"))
	itr := NewInteractor(presenter, &im.FakeImageStore{Return: existingImage})
	itr.Execute(Request{ImageId: existingImage.Id, Collection: existingImage.Collection.Name})
	got := presenter.Got
	if !presenter.GotSuccess {
		t.Fatalf("expected to get success")
	}
	if !(got.Id == existingImage.Id) || !(got.Collection == existingImage.Collection.Name) || !(len(got.Labels) == 1) {
		t.Fatalf("expected to get image id: %v, collection %v, num. labels %v, got %v, %v, %v",
			existingImage.Id, existingImage.Collection.Name, len(existingImage.Labels),
			got.Id, got.Collection, len(got.Labels))

	}
}
