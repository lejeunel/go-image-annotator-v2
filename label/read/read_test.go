package read

import (
	c "github.com/lejeunel/go-image-annotator-v2/label"
	"testing"
)

func TestReadLabel(t *testing.T) {
	name := "my-label"
	desc := "a-description"
	repo := &FakeReadRepo{Label: c.Label{Name: name, Description: desc}}
	presenter := &FakeReadPresenter{}
	itr := NewReadInteractor(repo, presenter)
	req := ReadRequest{Name: name}
	want := ReadResponse{Name: name, Description: desc}
	itr.Execute(req)
	if presenter.Got != want {
		t.Fatalf("expected %v, got %v", want, presenter.Got)
	}
}

func TestReadNonExistingLabelShouldFail(t *testing.T) {
	repo := &FakeReadRepo{Label: c.Label{Name: "my-label", Description: "a-description"}}
	presenter := &FakeReadPresenter{}
	itr := NewReadInteractor(repo, presenter)
	req := ReadRequest{Name: "non-existing-label"}
	itr.Execute(req)
	if !presenter.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeReadPresenter{}
	itr := NewReadInteractor(&FakeInternalErrReadRepo{}, presenter)
	itr.Execute(ReadRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
