package ingest

import (
	"testing"

	a "github.com/lejeunel/go-image-annotator-v2/domain/artefact"
	i "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingCollectionShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{MissingCollection: true}, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{})
	if !presenter.GotCollectionNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected not found error")
	}
}

func TestInvalidImageDataShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{Err: e.ErrInternal}})
	if !presenter.GotInvalidImageDataErr || presenter.GotSuccess {
		t.Fatal("expected invalid data error")
	}
}

func TestHandleInternalErrorOnCollectionExistsCheck(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnFindCollection: true, Err: e.ErrInternal}, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleArtefactRepoError(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{Err: e.ErrInternal}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected invalid data error")
	}
}

func TestNonExistingLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{MissingLabel: true}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label"}})
	if !presenter.GotLabelNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected label not found error, but go none")
	}
}

func TestHandleLabelExistsInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{ErrOnLabelExists: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label"}})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleIngestionInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{ErrOnIngest: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleAddLabelInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{ErrOnAddLabel: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label"}})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleValidationErrorOnAddLabel(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label", "a-label"}})
	if !presenter.GotValidationErr || presenter.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestAddImageDuplicateHashShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{HashAlreadyExists: true}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}})
	if !presenter.GotDuplicateImage || presenter.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestHandleDuplicateHashInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{ErrOnFindHash: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAddImageWithTwoLabels(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label", "another-label"}})
	if !presenter.GotSuccess || repo.NumLabelsAdded != 2 {
		t.Fatalf("expected to add two labels, but got %v", repo.NumLabelsAdded)
	}
}

func TestNonExistingBBoxLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{MissingLabel: true}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}, BoundingBoxes: []BoundingBoxRequest{{Label: "a-label"}}})
	if !presenter.GotLabelNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected label not found error, but go none")
	}
}

func TestHandleBoundingBoxValidationError(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: -2, Height: -4}}})
	if !presenter.GotValidationErr || presenter.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestHandleAddBoundingBoxInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{ErrOnAddBoundingBox: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: 2, Height: 4}}})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAddBoundingBoxToImage(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &a.FakeArtefactRepo{}, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: 2, Height: 4}}})
	if !presenter.GotSuccess || repo.NumBoundingboxesAdded != 1 {
		t.Fatalf("expected to add one bounding box to repo, but got %v", repo.NumBoundingboxesAdded)
	}
}

func TestInternalErrOnAddLabelMustDeleteImage(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{ErrOnAddLabel: true, Err: e.ErrInternal}
	artefactRepo := &a.FakeArtefactRepo{}
	itr := NewInteractor(repo, artefactRepo, presenter)
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label"}})
	if repo.NumDeletedImages != 1 || artefactRepo.NumDeletedImages != 1 {
		t.Fatal("expected to delete image")
	}
}
