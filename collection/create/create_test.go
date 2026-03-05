package create

import (
	"testing"
)

func TestCreateCollection(t *testing.T) {
	presenter := &FakeCreateCollectionPresenter{}
	repo := &FakeCreateCollectionRepo{}
	itr := NewCreateCollectionInteractor(repo, presenter)
	name := "a-name"
	desc := "a-description"
	req := CreateCollectionRequest{Name: name, Description: desc}
	want := CreateCollectionResponse{Name: name, Description: desc}
	itr.Execute(req)
	if presenter.Got != want {
		t.Fatalf("expected %v, got %v", want, presenter.Got)
	}
	if repo.Got != req {
		t.Fatalf("expected %v, got %v", req, repo.Got)

	}
}

func TestCreateCollectionWithDuplicateNameShouldFail(t *testing.T) {
	name := "my-collection"
	presenter := &FakeCreateCollectionPresenter{}
	itr := NewCreateCollectionInteractor(&FakeCreateCollectionRepo{Names: []string{name}}, presenter)
	itr.Execute(CreateCollectionRequest{Name: name})
	if !presenter.GotDuplicationErr {
		t.Fatal("expected duplication error, but go none")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeCreateCollectionPresenter{}
	itr := NewCreateCollectionInteractor(&FakeInternalErrCreateCollectionRepo{}, presenter)
	itr.Execute(CreateCollectionRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
