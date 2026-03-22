package list

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestListLabel(t *testing.T) {
	repo := &FakeRepo{}
	presenter := &FakePresenter{}
	itr := NewInteractor(repo, presenter)
	itr.Execute(Request{PageSize: 1, Page: 1})
	if !repo.ReturnedSomething {
		t.Fatal("expected repository to return something")
	}
	if !presenter.GotSuccess {
		t.Fatal("expected success")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeErrRepo{e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatalf("expected to get no success")
	}
}
