package delete

import (
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestDeleteLabelWithAssociatedResourcesShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{IsUsed_: true}, presenter)
	itr.Execute(Request{})
	if !presenter.GotDependencyErr {
		t.Fatal("expected dependency error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalErrOnIsUsed(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Err: e.ErrInternal, ErrOnIsUsed: true}, presenter)
	itr.Execute(Request{})
	if presenter.GotSuccess || !presenter.GotInternalErr {
		t.Fatal("expected internal error")
	}
}

func TestHandleInternalErrOnExists(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Err: e.ErrInternal, ErrOnExists: true}, presenter)
	itr.Execute(Request{})
	if presenter.GotSuccess || !presenter.GotInternalErr {
		t.Fatal("expected internal error")
	}
}

func TestDeletingMissingLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{IsMissing: true}, presenter)
	itr.Execute(Request{})
	if presenter.GotSuccess || !presenter.GotNotFoundErr {
		t.Fatal("expected not found error")
	}
}

func TestDeleteLabel(t *testing.T) {

	name := "my-collection"
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, presenter)
	itr.Execute(Request{Name: name})
	if !presenter.GotSuccess {
		t.Fatal("expected success, but did not")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Err: e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
