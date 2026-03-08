package delete

import (
	"testing"
)

func TestDeleteLabelWithAssociatedResourcesShouldFail(t *testing.T) {

	name := "my-collection"
	presenter := &FakeDeletePresenter{}
	itr := NewDeleteLabelInteractor(&FakeDeleteRepo{ArePopulated: []string{name}}, presenter)
	itr.Execute(DeleteRequest{Name: name})
	if !presenter.GotDependencyErr {
		t.Fatal("expected dependency error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestDeleteLabel(t *testing.T) {

	name := "my-collection"
	presenter := &FakeDeletePresenter{}
	itr := NewDeleteLabelInteractor(&FakeDeleteRepo{}, presenter)
	itr.Execute(DeleteRequest{Name: name})
	if !presenter.GotSuccess {
		t.Fatal("expected success, but did not")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeDeletePresenter{}
	itr := NewDeleteLabelInteractor(&FakeInternalErrDeleteRepo{}, presenter)
	itr.Execute(DeleteRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
