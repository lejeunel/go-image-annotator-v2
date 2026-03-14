package ingest

import (
	"io"
	"testing"

	a "github.com/lejeunel/go-image-annotator-v2/domain/artefact"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type FakeReader struct {
	Data   []byte
	Offset int
	Err    error
}

func (r FakeReader) Read(b []byte) (int, error) {
	if r.Err != nil {
		return 0, r.Err
	}
	if r.Offset >= len(r.Data) {
		return 0, io.EOF
	}
	n := copy(b, r.Data[r.Offset:])
	r.Offset += n
	return n, nil
}

func TestIngestInNonExistingCollectionShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{CollectionExists_: false}, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{})
	if !presenter.GotCollectionNotFoundErr {
		t.Fatal("expected not found error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestIngestInvalidImageDataShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{CollectionExists_: true, LabelExists_: true}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &FakeReader{Err: e.ErrInternal}})
	if !presenter.GotInvalidImageDataErr {
		t.Fatal("expected invalid data error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalErrorOnCollectionExistsCheck(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnFindCollection: true, Err: e.ErrInternal}, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}

func TestHandleArtefactRepoError(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{CollectionExists_: true, LabelExists_: true}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{Err: e.ErrInternal}, presenter)
	itr.Execute(Request{Reader: FakeReader{}})
	if !presenter.GotInternalErr {
		t.Fatal("expected invalid data error")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestIngestImageWithNonExistingLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{LabelExists_: false, CollectionExists_: true}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Labels: []string{"a-label"}})
	if !presenter.GotLabelNotFoundErr {
		t.Fatal("expected label not found error, but go none")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}

}

func TestHandleLabelExistsInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{CollectionExists_: true, ErrOnLabelExists: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Labels: []string{"a-label"}})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleIngestionInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{CollectionExists_: true, LabelExists_: true, ErrOnIngest: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: FakeReader{}})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error")
	}
	if presenter.GotSuccess {
		t.Fatal("expected no success")
	}
}
