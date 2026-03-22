package delete

import (
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestDeleteLabelWithAssociatedResourcesShouldFail(t *testing.T) {

	name := "my-collection"
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Used: []string{name}}, presenter)
	itr.Execute(Request{Name: name})
	if !presenter.GotDependencyErr {
		t.Fatal("expected dependency error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
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
	itr := NewInteractor(&FakeErrRepo{e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
