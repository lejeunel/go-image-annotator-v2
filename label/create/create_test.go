package create

import (
	"testing"
)

func TestCreateLabel(t *testing.T) {
	presenter := &FakeCreatePresenter{}
	repo := &FakeCreateRepo{}
	itr := NewCreateLabelInteractor(repo, presenter)
	name := "a-name"
	desc := "a-description"
	req := CreateLabelRequest{Name: name, Description: desc}
	wantp := CreateLabelResponse{Name: name, Description: desc}
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
	itr := NewCreateLabelInteractor(&FakeCreateRepo{Names: []string{name}}, presenter)
	itr.Execute(CreateLabelRequest{Name: name})
	if !presenter.GotDuplicationErr {
		t.Fatal("expected duplication error, but go none")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeCreatePresenter{}
	itr := NewCreateLabelInteractor(&FakeInternalErrCreateRepo{}, presenter)
	itr.Execute(CreateLabelRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
