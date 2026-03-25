package assign_label

import (
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestHandleNotFoundErrOnImageRetrieval(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, presenter, &im.FakeImageStore{Err: e.ErrNotFound})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected not found error")
	}
}

func TestHandleInternalErrOnImageRetrieval(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, presenter, &im.FakeImageStore{Err: e.ErrInternal})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAssignLabelToImage(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, presenter, &im.FakeImageStore{})
	imageId := im.NewImageId()
	label := "a-label"
	collection := "a-collection"
	req := Request{ImageId: imageId, Collection: collection, Label: label}
	itr.Execute(req)
	resp := presenter.Got
	if !presenter.GotSuccess || !(resp.Label == req.Label) || !(resp.Collection == req.Collection) || !(resp.ImageId == req.ImageId) {
		t.Fatalf("expected response does not match, got request %+v and response %+v", req, resp)
	}
}
