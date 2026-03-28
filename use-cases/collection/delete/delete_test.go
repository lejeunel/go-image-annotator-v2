package delete

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestDeleteNonExistingCollectionShouldFail(t *testing.T) {

	name := "my-collection"
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{})
	itr.Execute(DeleteRequest{Name: name}, p)
	if !p.GotNotFoundErr {
		t.Fatal("expected not found error, but got none")
	}
	if p.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestDeleteCollectionWithAssociatedResourcesShouldFail(t *testing.T) {

	name := "my-collection"
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Collections: []string{name}, ArePopulated: []string{name}})
	itr.Execute(DeleteRequest{Name: name}, p)
	if !p.GotDependencyErr {
		t.Fatal("expected dependency error, but got none")
	}
	if p.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestDeleteCollection(t *testing.T) {

	name := "my-collection"
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Collections: []string{name}})
	itr.Execute(DeleteRequest{Name: name}, p)
	if !p.GotSuccess {
		t.Fatal("expected success, but did not")
	}
}

func TestHandleInternalError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeErrRepo{e.ErrInternal})
	itr.Execute(DeleteRequest{}, p)
	if !p.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}
