package delete

import (
	"testing"
)

func TestDeleteLabelWithAssociatedResourcesShouldFail(t *testing.T) {

	name := "my-collection"
	presenter := &FakeDeleteLabelPresenter{}
	itr := NewDeleteLabelInteractor(&FakeDeleteLabelRepo{ArePopulated: []string{name}}, presenter)
	itr.Execute(DeleteLabelRequest{Name: name})
	if !presenter.GotDependencyErr {
		t.Fatal("expected dependency error, but got none")
	}
}

func TestDeleteLabel(t *testing.T) {

	name := "my-collection"
	presenter := &FakeDeleteLabelPresenter{}
	itr := NewDeleteLabelInteractor(&FakeDeleteLabelRepo{}, presenter)
	itr.Execute(DeleteLabelRequest{Name: name})
	if !presenter.GotSuccess {
		t.Fatal("expected success, but did not")
	}
}

func TestHandleInternalError(t *testing.T) {
	presenter := &FakeDeleteLabelPresenter{}
	itr := NewDeleteLabelInteractor(&FakeInternalErrDeleteLabelRepo{}, presenter)
	itr.Execute(DeleteLabelRequest{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
