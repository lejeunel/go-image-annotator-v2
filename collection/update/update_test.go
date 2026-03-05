package update

import (
	"testing"
)

func TestUpdateNonExistingCollectionShouldFail(t *testing.T) {

	presenter := &FakeUpdateCollectionPresenter{}
	non_existing_name := "non-existing-name"
	itr := NewUpdateCollectionInteractor(&FakeUpdateCollectionRepo{}, presenter)
	itr.Execute(UpdateCollectionRequest{Name: non_existing_name, NewName: "new-name"})
	if !presenter.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
}

func TestUpdateCollection(t *testing.T) {
	name := "name"
	new_name := "updated-name"
	new_description := "updated-description"

	presenter := &FakeUpdateCollectionPresenter{}
	repo := &FakeUpdateCollectionRepo{Names: []string{name}}
	itr := NewUpdateCollectionInteractor(repo, presenter)
	req := UpdateCollectionRequest{Name: name, NewName: new_name, NewDescription: new_description}
	want := UpdateCollectionResponse{Name: new_name, Description: new_description}
	itr.Execute(req)
	if presenter.Got != want {
		t.Fatalf("expected %v, got %v", want, presenter.Got)

	}
	if repo.Got != req {
		t.Fatalf("expected %v, got %v", req, repo.Got)

	}
}

func TestUpdateCollectionWithNameAlreadyTakenShouldFail(t *testing.T) {

	presenter := &FakeUpdateCollectionPresenter{}
	name := "name"
	existing_name := "existing-name"
	itr := NewUpdateCollectionInteractor(&FakeUpdateCollectionRepo{Names: []string{name, existing_name}}, presenter)
	itr.Execute(UpdateCollectionRequest{Name: name, NewName: existing_name})
	if !presenter.GotDuplicationErr {
		t.Fatal("expected duplication error, but got none")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeUpdateCollectionPresenter{}
	itr := NewUpdateCollectionInteractor(&FakeInternalErrUpdateCollectionRepo{}, presenter)
	itr.Execute(UpdateCollectionRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
