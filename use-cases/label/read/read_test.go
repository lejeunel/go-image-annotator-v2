package read

import (
	l "github.com/lejeunel/go-image-annotator-v2/entities/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestReadLabel(t *testing.T) {
	name := "my-label"
	desc := "a-description"
	repo := &FakeRepo{Label: l.Label{Name: name, Description: desc}}
	p := &FakePresenter{}
	itr := NewInteractor(repo)
	req := Request{Name: name}
	want := Response{Name: name, Description: desc}
	itr.Execute(req, p)
	if p.Got != want {
		t.Fatalf("expected %v, got %v", want, p.Got)
	}
}

func TestReadNonExistingLabelShouldFail(t *testing.T) {
	repo := &FakeRepo{Label: l.Label{Name: "my-label", Description: "a-description"}}
	p := &FakePresenter{}
	itr := NewInteractor(repo)
	req := Request{Name: "non-existing-label"}
	itr.Execute(req, p)
	if !p.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
	if p.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Err: e.ErrInternal})
	itr.Execute(Request{}, p)
	if !p.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
