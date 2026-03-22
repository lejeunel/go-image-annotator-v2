package list

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestListCollection(t *testing.T) {
	repo := &FakeRepo{}
	presenter := &FakeListPresenter{}
	itr := NewInteractor(repo, presenter)
	req := Request{PageSize: 1, Page: 1}
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
	itr := NewInteractor(&FakeErrListRepo{e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatalf("expected to get no success")
	}
}
