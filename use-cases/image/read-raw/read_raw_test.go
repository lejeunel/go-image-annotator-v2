package read_raw

import (
	"bytes"
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingImageShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Err: e.ErrNotFound})
	itr.Execute(Request{ImageId: im.NewImageId()}, p)
	if !p.GotNotFoundErr || p.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestHandleInternalErr(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Err: e.ErrInternal})
	itr.Execute(Request{ImageId: im.NewImageId()}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestReadOriginalBytes(t *testing.T) {
	p := &FakePresenter{}
	data := []byte("test-data")
	itr := NewInteractor(&FakeRepo{Data: data})
	itr.Execute(Request{ImageId: im.NewImageId()}, p)
	gotData := p.Got.Data
	if !p.GotSuccess || !bytes.Equal(gotData, data) {
		t.Fatalf("expected to retrieve input data (%v), got %v", data, gotData)
	}
}
