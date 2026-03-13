package ingest

import (
	"bytes"
	"strings"
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type FakeFailingReader struct{}

func (r *FakeFailingReader) Read(b []byte) (int, error) {
	return 0, e.ErrValidation
}

func TestIngestInNonExistingCollectionShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, &FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Collection: "non-existing-collection", Reader: bytes.NewReader([]byte{})})
	if !presenter.GotCollectionNotFoundErr {
		t.Fatal("expected not found error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestIngestInvalidImageDataShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{Collections: []string{"a-collection"}}
	itr := NewInteractor(repo, &FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Collection: "a-collection", Reader: &FakeFailingReader{}})
	if !presenter.GotInvalidImageDataErr {
		t.Fatal("expected invalid data error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalErrorOnCollectionExistsCheck(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeCollectionExistsErrRepo{e.ErrInternal}, &FakeArtefactRepo{}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}

func TestHandleArtefactRepoError(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{Collections: []string{"a-collection"}}
	itr := NewInteractor(repo, &FakeErrArtefactRepo{e.ErrInternal}, presenter)
	itr.Execute(Request{Collection: "a-collection", Reader: strings.NewReader("dummy-data")})
	if !presenter.GotInternalErr {
		t.Fatal("expected invalid data error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestIngestImageWithNonExistingLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{Collections: []string{"a-collection"}, Labels: []string{"a-label"}}
	itr := NewInteractor(repo, &FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Collection: "a-collection", Labels: []string{"non-existing-label"}})
	if !presenter.GotLabelNotFoundErr {
		t.Fatal("expected label not found error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}

}

func TestHandleLabelExistsInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeLabelExistsErrRepo{e.ErrInternal}
	itr := NewInteractor(repo, &FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Collection: "a-collection", Labels: []string{"a-label"}})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}

}
