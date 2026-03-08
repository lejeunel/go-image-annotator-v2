package update

import (
	"testing"
)

func TestUpdateNonExistingCollectionShouldFail(t *testing.T) {

	presenter := &FakeUpdatePresenter{}
	non_existing_name := "non-existing-name"
	itr := NewUpdateCollectionInteractor(&FakeUpdateRepo{}, presenter)
	itr.Execute(UpdateRequest{Name: non_existing_name, NewName: "new-name"})
	if !presenter.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestUpdateCollection(t *testing.T) {
	name := "name"
	new_name := "updated-name"
	new_description := "updated-description"

	presenter := &FakeUpdatePresenter{}
	repo := &FakeUpdateRepo{Names: []string{name}}
	itr := NewUpdateCollectionInteractor(repo, presenter)
	req := UpdateRequest{Name: name, NewName: new_name, NewDescription: new_description}
	wantr := UpdateModel{Name: name, NewName: new_name, NewDescription: new_description}
	wantp := UpdateResponse{Name: new_name, Description: new_description}
	itr.Execute(req)
	if presenter.Got != wantp {
		t.Fatalf("expected %v, got %v", wantp, presenter.Got)

	}
	if repo.Got != wantr {
		t.Fatalf("expected %v, got %v", req, repo.Got)

	}
}

func TestUpdateCollectionWithNameAlreadyTakenShouldFail(t *testing.T) {

	presenter := &FakeUpdatePresenter{}
	name := "name"
	existing_name := "existing-name"
	itr := NewUpdateCollectionInteractor(&FakeUpdateRepo{Names: []string{name, existing_name}}, presenter)
	itr.Execute(UpdateRequest{Name: name, NewName: existing_name})
	if !presenter.GotDuplicationErr {
		t.Fatal("expected duplication error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeUpdatePresenter{}
	itr := NewUpdateCollectionInteractor(&FakeInternalErrUpdateRepo{}, presenter)
	itr.Execute(UpdateRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
