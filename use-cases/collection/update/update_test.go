package update

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestUpdateNonExistingCollectionShouldFail(t *testing.T) {

	presenter := &FakePresenter{}
	non_existing_name := "non-existing-name"
	itr := NewUpdateCollectionInteractor(&FakeRepo{}, presenter)
	itr.Execute(Request{Name: non_existing_name, NewName: "new-name"})
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

	presenter := &FakePresenter{}
	repo := &FakeRepo{Names: []string{name}}
	itr := NewUpdateCollectionInteractor(repo, presenter)
	req := Request{Name: name, NewName: new_name, NewDescription: new_description}
	wantr := Model{Name: name, NewName: new_name, NewDescription: new_description}
	wantp := Response{Name: new_name, Description: new_description}
	itr.Execute(req)
	if presenter.Got != wantp {
		t.Fatalf("expected %v, got %v", wantp, presenter.Got)

	}
	if repo.Got != wantr {
		t.Fatalf("expected %v, got %v", req, repo.Got)

	}
}

func TestUpdateCollectionWithNameAlreadyTakenShouldFail(t *testing.T) {

	presenter := &FakePresenter{}
	name := "name"
	existing_name := "existing-name"
	itr := NewUpdateCollectionInteractor(&FakeRepo{Names: []string{name, existing_name}}, presenter)
	itr.Execute(Request{Name: name, NewName: existing_name})
	if !presenter.GotDuplicationErr {
		t.Fatal("expected duplication error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewUpdateCollectionInteractor(&FakeErrRepo{e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
