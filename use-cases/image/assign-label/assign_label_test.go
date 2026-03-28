package assign_label

import (
	"testing"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestHandleNotFoundErrOnImageRetrieval(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, &im.FakeImageStore{Err: e.ErrNotFound})
	itr.Execute(Request{}, p)
	if !p.GotNotFoundErr || p.GotSuccess {
		t.Fatal("expected not found error")
	}
}

func TestHandleInternalErrOnImageRetrieval(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, &im.FakeImageStore{Err: e.ErrInternal})
	itr.Execute(Request{}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAssignNonExistingLabelShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{MissingLabel: true}, &im.FakeImageStore{})
	itr.Execute(Request{}, p)
	if !p.GotNotFoundErr || p.GotSuccess {
		t.Fatal("expected not found error")
	}
}
func TestInternalErrOnFindLabelShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnFindLabel: true, Err: e.ErrInternal}, &im.FakeImageStore{})
	itr.Execute(Request{}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAssignLabelToImage(t *testing.T) {
	p := &FakePresenter{}
	collection := clc.NewCollection(clc.NewCollectionId(), "a-collection")
	image := im.NewImage(im.NewImageId(), *collection)
	label := lbl.NewLabel(lbl.NewLabelId(), "al-label")
	req := Request{ImageId: image.Id, Collection: collection.Name, Label: label.Name}
	repo := &FakeRepo{ReturnLabel: *label}
	itr := NewInteractor(repo, &im.FakeImageStore{Return: image})
	itr.Execute(req, p)
	resp := p.Got
	if !p.GotSuccess || !(resp.Label == req.Label) || !(resp.Collection == req.Collection) || !(resp.ImageId == req.ImageId) {
		t.Fatalf("expected response does not match, got request %+v and response %+v", req, resp)
	}
	if (repo.AddedLabelId != label.Id) || (repo.AddedOnImageId != image.Id) || (repo.AddedOnCollectionId != collection.Id) {
		t.Fatalf("expected to add label %v on image %v and collection %v, got %v, %v, %v",
			label, image.Id, image.Collection.Id,
			repo.AddedLabelId, repo.AddedOnImageId, repo.AddedOnCollectionId)
	}
}
