package read

import (
	"testing"
)

func TestReadCollection(t *testing.T) {
	name := "my-collection"
	desc := "a-description"
	repo := &FakeReadCollectionRepo{Collection: Collection{Name: name, Description: desc}}
	presenter := &FakeReadCollectionPresenter{}
	itr := NewReadCollectionInteractor(repo, presenter)
	req := ReadCollectionRequest{Name: name}
	want := ReadCollectionResponse{Name: name, Description: desc}
	itr.Execute(req)
	if presenter.Got != want {
		t.Fatalf("expected %v, got %v", want, presenter.Got)
	}
}

func TestReadNonExistingCollectionShouldFail(t *testing.T) {
	repo := &FakeReadCollectionRepo{Collection: Collection{Name: "my-collection", Description: "a-description"}}
	presenter := &FakeReadCollectionPresenter{}
	itr := NewReadCollectionInteractor(repo, presenter)
	req := ReadCollectionRequest{Name: "non-existing-collection"}
	itr.Execute(req)
	if !presenter.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeReadCollectionPresenter{}
	itr := NewReadCollectionInteractor(&FakeInternalErrReadCollectionRepo{}, presenter)
	itr.Execute(ReadCollectionRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
