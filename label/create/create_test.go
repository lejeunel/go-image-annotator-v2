package create

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
	"testing"
)

func TestCreateLabel(t *testing.T) {
	presenter := &FakeCreatePresenter{}
	repo := &FakeCreateRepo{}
	itr := NewCreateLabelInteractor(repo, &v.FakeValidNameValidator{}, presenter)
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

func TestCreateLabelWithDuplicateNameShouldFail(t *testing.T) {
	name := "my-label"
	presenter := &FakeCreatePresenter{}
	itr := NewCreateLabelInteractor(&FakeCreateRepo{Names: []string{name}}, &v.FakeValidNameValidator{}, presenter)
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
	itr := NewCreateLabelInteractor(&FakeErrCreateRepo{e.ErrInternal}, &v.FakeValidNameValidator{}, presenter)
	itr.Execute(CreateRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}

func TestCreateLabelWithInvalidNameShouldFail(t *testing.T) {
	presenter := &FakeCreatePresenter{}
	itr := NewCreateLabelInteractor(&FakeCreateRepo{}, &v.FakeInvalidNameValidator{}, presenter)
	itr.Execute(CreateRequest{Name: "invalid-name"})
	if !presenter.GotValidationErr {
		t.Fatal("expected validation error, but go none")
	}
}
