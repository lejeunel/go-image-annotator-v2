package delete

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestDeleteNonExistingCollectionShouldFail(t *testing.T) {

	name := "my-collection"
	presenter := &FakePresenter{}
	itr := NewDeleteInteractor(&FakeRepo{}, presenter)
	itr.Execute(DeleteRequest{Name: name})
	if !presenter.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestDeleteCollectionWithAssociatedResourcesShouldFail(t *testing.T) {

	name := "my-collection"
	presenter := &FakePresenter{}
	itr := NewDeleteInteractor(&FakeRepo{Collections: []string{name}, ArePopulated: []string{name}}, presenter)
	itr.Execute(DeleteRequest{Name: name})
	if !presenter.GotDependencyErr {
		t.Fatal("expected dependency error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestDeleteCollection(t *testing.T) {

	name := "my-collection"
	presenter := &FakePresenter{}
	itr := NewDeleteInteractor(&FakeRepo{Collections: []string{name}}, presenter)
	itr.Execute(DeleteRequest{Name: name})
	if !presenter.GotSuccess {
		t.Fatal("expected success, but did not")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewDeleteInteractor(&FakeErrRepo{e.ErrInternal}, presenter)
	itr.Execute(DeleteRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
