package create

import (
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
)

func TestCreateCollection(t *testing.T) {
	presenter := &FakeCreatePresenter{}
	repo := &FakeCreateRepo{}
	itr := NewCreateInteractor(repo, &v.FakeValidNameValidator{}, presenter)
	name := "a-name"
	desc := "a-description"
	req := CreateRequest{Name: name, Description: desc}
	wantp := CreateResponse{Name: name, Description: desc}
	wantr := CreateModel{Name: name, Description: desc}
	itr.Execute(req)
	if presenter.Got != wantp {
		t.Fatalf("expected %v, got %v", wantp, presenter.Got)
	}
	if repo.Got != wantr {
		t.Fatalf("expected %v, got %v", wantr, repo.Got)
	}
}

func TestCreateCollectionWithDuplicateNameShouldFail(t *testing.T) {
	name := "my-collection"
	presenter := &FakeCreatePresenter{}
	itr := NewCreateInteractor(&FakeCreateRepo{Names: []string{name}}, &v.FakeValidNameValidator{}, presenter)
	itr.Execute(CreateRequest{Name: name})
	if !presenter.GotDuplicationErr {
		t.Fatal("expected duplication error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeCreatePresenter{}
	itr := NewCreateInteractor(&FakeErrCreateRepo{e.ErrInternal}, &v.FakeValidNameValidator{}, presenter)
	itr.Execute(CreateRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}

func TestCreateCollectionWithInvalidNameShouldFail(t *testing.T) {
	name := "my-collection%/"
	presenter := &FakeCreatePresenter{}
	validator := &v.FakeInvalidNameValidator{}
	itr := NewCreateInteractor(&FakeCreateRepo{Names: []string{name}}, validator, presenter)
	itr.Execute(CreateRequest{Name: name})
	if !presenter.GotValidationErr {
		t.Fatal("expected validation error, but go none")
	}
}
