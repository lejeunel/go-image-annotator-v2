package delete

import (
	"testing"
)

func TestDeleteCollectionWithAssociatedResourcesShouldFail(t *testing.T) {

	name := "my-collection"
	presenter := &FakeDeletePresenter{}
	itr := NewDeleteInteractor(&FakeDeleteRepo{ArePopulated: []string{name}}, presenter)
	itr.Execute(DeleteRequest{Name: name})
	if !presenter.GotDependencyErr {
		t.Fatal("expected dependency error, but got none")
	}
}

func TestDeleteCollection(t *testing.T) {

	name := "my-collection"
	presenter := &FakeDeletePresenter{}
	itr := NewDeleteInteractor(&FakeDeleteRepo{}, presenter)
	itr.Execute(DeleteRequest{Name: name})
	if !presenter.GotSuccess {
		t.Fatal("expected success, but did not")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeDeletePresenter{}
	itr := NewDeleteInteractor(&FakeInternalErrDeleteRepo{}, presenter)
	itr.Execute(DeleteRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
