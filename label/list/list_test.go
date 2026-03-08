package read

import (
	"testing"
)

func TestListLabel(t *testing.T) {
	repo := &FakeListRepo{}
	presenter := &FakeListPresenter{}
	itr := NewListInteractor(repo, presenter)
	req := ListRequest{PageSize: 1, Page: 1}
	itr.Execute(req)
	if !repo.ReturnedSomething {
		t.Fatal("expected repository to return something")
	}
	if !presenter.GotSuccess {
		t.Fatal("expected success")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeListPresenter{}
	itr := NewListInteractor(&FakeInternalErrListRepo{}, presenter)
	itr.Execute(ListRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatalf("expected to get no success")
	}
}
