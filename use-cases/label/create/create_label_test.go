package create

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
	"testing"
)

func TestCreateLabelWithDuplicateNameShouldFail(t *testing.T) {
	name := "my-label"
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Names: []string{name}}, &v.FakeNameValidator{}, presenter)
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
	itr := NewInteractor(&FakeRepo{Err: e.ErrInternal}, &v.FakeNameValidator{}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}

func TestCreateLabelWithInvalidNameShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, &v.FakeNameValidator{Err: e.ErrValidation}, presenter)
	itr.Execute(Request{Name: "invalid-name"})
	if !presenter.GotValidationErr {
		t.Fatal("expected validation error, but go none")
	}
}

func TestCreateLabel(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &v.FakeNameValidator{}, presenter)
	name := "a-name"
	desc := "a-description"
	req := Request{Name: name, Description: desc}
	wantp := Response{Name: name, Description: desc}
	itr.Execute(req)
	if presenter.Got != wantp {
		t.Fatalf("expected %v, got %v", wantp, presenter.Got)
	}
	if repo.Got.Name != name || repo.Got.Description != desc || repo.Got.Id.IsNil() {
		t.Fatalf("expected to create label with name %v, description %v, and non-nil id, got %v, %v, %v",
			name, desc, repo.Got.Name, repo.Got.Description, repo.Got.Id)

	}
}
