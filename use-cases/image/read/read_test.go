package read

import (
	"testing"

	st "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func TestHandleNotFoundError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&st.FakeImageStore{Err: e.ErrNotFound})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"}, p)
	if !p.GotNotFoundErr || p.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestHandleInternalError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&st.FakeImageStore{Err: e.ErrInternal})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestFindImage(t *testing.T) {
	p := &FakePresenter{}
	existingImage := im.NewImage(im.NewImageId(), *clc.NewCollection(clc.NewCollectionId(), "a-collection"))
	itr := NewInteractor(&st.FakeImageStore{Return: existingImage})
	itr.Execute(Request{ImageId: existingImage.Id, Collection: existingImage.Collection.Name}, p)
	got := p.Got
	if !p.GotSuccess {
		t.Fatalf("expected to get success")
	}
	if !(got.Id == existingImage.Id) || !(got.Collection.Name == existingImage.Collection.Name) {
		t.Fatalf("expected to get image id: %v, collection %v, got %v, %v",
			existingImage.Id, existingImage.Collection.Name,
			got.Id, got.Collection)

	}
}
