package create

import (
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	v "github.com/lejeunel/go-image-annotator-v2/shared/validation"
)

func TestCreateCollectionWithDuplicateNameShouldFail(t *testing.T) {
	name := "my-collection"
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Names: []string{name}}, &v.FakeNameValidator{})
	itr.Execute(Request{Name: name}, p)
	if !p.GotDuplicationErr {
		t.Fatal("expected duplication error, but go none")
	}
	if p.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Err: e.ErrInternal}, &v.FakeNameValidator{})
	itr.Execute(Request{}, p)
	if !p.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}

func TestCreateCollectionWithInvalidNameShouldFail(t *testing.T) {
	name := "my-collection%/"
	p := &FakePresenter{}
	validator := &v.FakeNameValidator{Err: e.ErrValidation}
	itr := NewInteractor(&FakeRepo{Names: []string{name}}, validator)
	itr.Execute(Request{Name: name}, p)
	if !p.GotValidationErr {
		t.Fatal("expected validation error, but go none")
	}
}

func TestCreateCollection(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &v.FakeNameValidator{})
	name := "a-name"
	desc := "a-description"
	req := Request{Name: name, Description: desc}
	wantp := Response{Name: name, Description: desc}
	itr.Execute(req, p)
	if p.Got != wantp {
		t.Fatalf("expected %v, got %v", wantp, p.Got)
	}
	if repo.Got.Name != name || repo.Got.Description != desc || repo.Got.Id.IsNil() {
		t.Fatalf("expected to create label with name %v, description %v, and non-nil id, got %v, %v, %v",
			name, desc, repo.Got.Name, repo.Got.Description, repo.Got.Id)

	}
}
