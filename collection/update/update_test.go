package update

import (
	"testing"
)

func TestUpdateCollection(t *testing.T) {

	presenter := &FakeUpdateCollectionPresenter{}
	repo := &FakeUpdateCollectionRepo{}
	itr := NewUpdateCollectionInteractor(repo, presenter)
	new_name := "updated-name"
	new_description := "updated-description"
	req := UpdateCollectionRequest{Name: new_name, Description: new_description}
	res := UpdateCollectionResponse{Name: new_name, Description: new_description}
	itr.Execute(req)
	if presenter.Got != res {
		t.Fatalf("expected %v, got %v", res, presenter.Got)

	}
	if repo.Got != req {
		t.Fatalf("expected %v, got %v", req, repo.Got)

	}
}

func TestUpdateCollectionWithExistingNameShouldFail(t *testing.T) {

	presenter := &FakeUpdateCollectionPresenter{}
	new_name := "updated-name"
	itr := NewUpdateCollectionInteractor(&FakeUpdateCollectionRepo{Names: []string{new_name}}, presenter)
	itr.Execute(UpdateCollectionRequest{Name: new_name})
	if !presenter.GotDuplicationErr {
		t.Fatal("expected duplication error, but go none")
	}
}
