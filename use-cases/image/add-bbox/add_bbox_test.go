package add_bbox

import (
	"testing"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingImageStoreResourceShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{Err: e.ErrNotFound}, &FakeRepo{})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestInternalErrInImageStoreShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{Err: e.ErrInternal}, &FakeRepo{})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected internal error")
	}
}

func TestInternalErrOnImageRetrievalShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{Err: e.ErrInternal}, &FakeRepo{})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected internal error")
	}
}

func TestNotFoundErrOnImageRetrievalShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{Err: e.ErrNotFound}, &FakeRepo{})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected not found error")
	}
}

func TestNotFoundErrOnFindLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{}, &FakeRepo{ErrOnFindLabel: true, Err: e.ErrNotFound})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected not found error")
	}
}

func TestInternalErrOnFindLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{}, &FakeRepo{ErrOnFindLabel: true, Err: e.ErrInternal})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected internal error")
	}
}

func TestValidationErrOnAddBoxShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{}, &FakeRepo{})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection", Label: "a-label",
		Xc: 1, Yc: 1, Width: -999, Height: 3})
	if !presenter.GotValidationErr || presenter.GotSuccess {
		t.Fatalf("expected validation error")
	}
}

func TestNotFoundErrOnAddBoxShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{}, &FakeRepo{ErrOnAdd: true, Err: e.ErrNotFound})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection", Label: "a-label",
		Xc: 1, Yc: 1, Width: 3, Height: 3})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected not found error")
	}
}

func TestInternalErrOnAddBoxShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &im.FakeImageStore{}, &FakeRepo{ErrOnAdd: true, Err: e.ErrInternal})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection", Label: "a-label",
		Xc: 1, Yc: 1, Width: 3, Height: 3})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected internal error")
	}
}

func TestAddBoundingBox(t *testing.T) {
	presenter := &FakePresenter{}
	repo := FakeRepo{}
	imageId := im.NewImageId()
	collection := clc.NewCollection("a-collection")
	image := im.NewImage(imageId, *collection)
	label := "a-label"
	x := float32(1.0)
	y := float32(1.0)
	width := float32(3.0)
	height := float32(3.0)
	itr := NewInteractor(presenter, &im.FakeImageStore{Return: image}, &repo)
	itr.Execute(Request{ImageId: imageId, Collection: "a-collection", Label: label,
		Xc: x, Yc: y, Width: width, Height: height})
	if !presenter.GotSuccess {
		t.Fatalf("expected success")
	}
	if repo.GotImageId != imageId {
		t.Fatalf("expected to store box on image %v, got %v", imageId.String(), repo.GotImageId.String())
	}
	if repo.GotCollectionId != collection.Id {
		t.Fatalf("expected to store box on collection %v, got %v", collection.Id.String(), repo.GotCollectionId.String())
	}
	if repo.GotBox.Label.Name != label {
		t.Fatalf("expected to store box with label %v, got %v", label, repo.GotBox.Label.Name)
	}
	if repo.GotBox.Xc != x {
		t.Fatalf("expected to store box with x %v, got %v", x, repo.GotBox.Xc)
	}
	if repo.GotBox.Yc != y {
		t.Fatalf("expected to store box with y %v, got %v", x, repo.GotBox.Yc)
	}
	if repo.GotBox.Width != width {
		t.Fatalf("expected to store box with width %v, got %v", width, repo.GotBox.Width)
	}
	if repo.GotBox.Height != height {
		t.Fatalf("expected to store box with height %v, got %v", height, repo.GotBox.Height)
	}

}
