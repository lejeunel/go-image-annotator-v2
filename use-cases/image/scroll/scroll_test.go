package scroll

import (
	"testing"

	st "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	// e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func TestErrOnCurrentImageShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&st.FakeImageStore{}, &FakeImageRepo{Err: e.ErrNotFound})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"}, p)
	if !p.GotErr || p.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestScrollNextByDefault(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeImageRepo{}
	itr := NewInteractor(&st.FakeImageStore{}, repo)
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"}, p)
	if !repo.ReturnedNext || !p.GotSuccess {
		t.Fatal("expected to retrieve next image by default")
	}
}

func TestScrollPrevious(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeImageRepo{}
	itr := NewInteractor(&st.FakeImageStore{}, repo)
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection",
		Direction: ScrollPrevious}, p)
	if !repo.ReturnedPrev || !p.GotSuccess {
		t.Fatal("expected to retrieve prev image")
	}
}

func TestHandleErrorOnImageStore(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&st.FakeImageStore{Err: e.ErrInternal}, &FakeImageRepo{})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"}, p)
	if !p.GotErr || p.GotSuccess {
		t.Fatalf("expected to get error")
	}
}

func TestReturnImage(t *testing.T) {
	p := &FakePresenter{}
	image := im.NewImage(im.NewImageId(), *clc.NewCollection(clc.NewCollectionId(), ""))
	itr := NewInteractor(&st.FakeImageStore{Return: image}, &FakeImageRepo{})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"}, p)
	if p.Got.Id != image.Id || !p.GotSuccess {
		t.Fatalf("expected to retrieve image")
	}
}
