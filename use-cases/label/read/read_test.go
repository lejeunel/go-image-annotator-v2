package read

import (
	l "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestReadLabel(t *testing.T) {
	name := "my-label"
	desc := "a-description"
	repo := &FakeRepo{Label: l.Label{Name: name, Description: desc}}
	presenter := &FakePresenter{}
	itr := NewReadInteractor(repo, presenter)
	req := Request{Name: name}
	want := Response{Name: name, Description: desc}
	itr.Execute(req)
	if presenter.Got != want {
		t.Fatalf("expected %v, got %v", want, presenter.Got)
	}
}

func TestReadNonExistingLabelShouldFail(t *testing.T) {
	repo := &FakeRepo{Label: l.Label{Name: "my-label", Description: "a-description"}}
	presenter := &FakePresenter{}
	itr := NewReadInteractor(repo, presenter)
	req := Request{Name: "non-existing-label"}
	itr.Execute(req)
	if !presenter.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewReadInteractor(&FakeErrRepo{e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
