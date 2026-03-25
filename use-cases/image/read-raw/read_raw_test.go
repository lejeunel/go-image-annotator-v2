package read_raw

import (
	"bytes"
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingImageShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{Err: e.ErrNotFound})
	itr.Execute(Request{ImageId: im.NewImageId()})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestHandleInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{Err: e.ErrInternal})
	itr.Execute(Request{ImageId: im.NewImageId()})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestReadOriginalBytes(t *testing.T) {
	presenter := &FakePresenter{}
	data := []byte("test-data")
	itr := NewInteractor(presenter, &FakeRepo{Data: data})
	itr.Execute(Request{ImageId: im.NewImageId()})
	gotData := presenter.Got.Data
	if !presenter.GotSuccess || !bytes.Equal(gotData, data) {
		t.Fatalf("expected to retrieve input data (%v), got %v", data, gotData)
	}
}
