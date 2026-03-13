package create

import (
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
)

func TestCreateCollection(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &v.FakeValidNameValidator{}, presenter)
	name := "a-name"
	desc := "a-description"
	req := Request{Name: name, Description: desc}
	wantp := Response{Name: name, Description: desc}
	wantr := Model{Name: name, Description: desc}
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
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Names: []string{name}}, &v.FakeValidNameValidator{}, presenter)
	itr.Execute(Request{Name: name})
	if !presenter.GotDuplicationErr {
		t.Fatal("expected duplication error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeErrRepo{e.ErrInternal}, &v.FakeValidNameValidator{}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}

func TestCreateCollectionWithInvalidNameShouldFail(t *testing.T) {
	name := "my-collection%/"
	presenter := &FakePresenter{}
	validator := &v.FakeInvalidNameValidator{}
	itr := NewInteractor(&FakeRepo{Names: []string{name}}, validator, presenter)
	itr.Execute(Request{Name: name})
	if !presenter.GotValidationErr {
		t.Fatal("expected validation error, but go none")
	}
}
