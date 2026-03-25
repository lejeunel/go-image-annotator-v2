package read

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	"testing"
)

func TestReadCollection(t *testing.T) {
	name := "my-collection"
	desc := "a-description"
	repo := &FakeReadRepo{Collection: clc.Collection{Name: name, Description: desc}}
	presenter := &FakeReadPresenter{}
	itr := NewInteractor(repo, presenter)
	req := Request{Name: name}
	want := Response{Name: name, Description: desc}
	itr.Execute(req)
	if presenter.Got != want {
		t.Fatalf("expected %v, got %v", want, presenter.Got)
	}
}

func TestReadNonExistingCollectionShouldFail(t *testing.T) {
	repo := &FakeReadRepo{Collection: clc.Collection{Name: "my-collection", Description: "a-description"}}
	presenter := &FakeReadPresenter{}
	itr := NewInteractor(repo, presenter)
	req := Request{Name: "non-existing-collection"}
	itr.Execute(req)
	if !presenter.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeReadPresenter{}
	itr := NewInteractor(&FakeInternalErrReadRepo{}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
