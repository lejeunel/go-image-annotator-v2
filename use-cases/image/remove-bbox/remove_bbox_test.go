package remove_bbox

import (
	"testing"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingBoxShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{Err: e.ErrNotFound})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestInternalErrShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{Err: e.ErrInternal})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected internal error")
	}
}

func TestRemoveBox(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(presenter, repo)
	annotationId := a.NewAnnotationId()
	itr.Execute(Request{Id: annotationId})
	if !presenter.GotSuccess || !(repo.Got == annotationId) {
		t.Fatalf("expected to remove annotation, got success %v and annotation id %v, wanted %v",
			presenter.GotSuccess, repo.Got.String(), annotationId.String())
	}
}
