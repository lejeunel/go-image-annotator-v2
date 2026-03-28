package ingest

import (
	"testing"

	i "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingCollectionShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{MissingCollection: true}, &i.FakeArtefactRepo{})
	itr.Execute(Request{}, p)
	if !p.GotCollectionNotFoundErr || p.GotSuccess {
		t.Fatal("expected not found error")
	}
}

func TestInvalidImageDataShouldFail(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{Err: e.ErrInternal}}, p)
	if !p.GotInvalidImageDataErr || p.GotSuccess {
		t.Fatal("expected invalid data error")
	}
}

func TestHandleInternalErrorOnCollectionExistsCheck(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnFindCollection: true, Err: e.ErrInternal}, &i.FakeArtefactRepo{})
	itr.Execute(Request{}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleArtefactRepoError(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{Err: e.ErrInternal})
	itr.Execute(Request{Reader: &i.FakeImageReader{}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected invalid data error")
	}
}

func TestNonExistingLabelShouldFail(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{MissingLabel: true}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if !p.GotLabelNotFoundErr || p.GotSuccess {
		t.Fatal("expected label not found error, but go none")
	}
}

func TestHandleLabelExistsInternalErr(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{ErrOnLabelExists: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleIngestionInternalErr(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{ErrOnIngest: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleAddLabelInternalErr(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{ErrOnAddLabel: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleValidationErrorOnAddLabel(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label", "a-label"}}, p)
	if !p.GotValidationErr || p.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestAddImageDuplicateHashShouldFail(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{HashAlreadyExists: true}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}}, p)
	if !p.GotDuplicateImage || p.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestHandleDuplicateHashInternalErr(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{ErrOnFindHash: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAddImageWithTwoLabels(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label", "another-label"}}, p)
	if !p.GotSuccess || repo.NumLabelsAdded != 2 {
		t.Fatalf("expected to add two labels, but got %v", repo.NumLabelsAdded)
	}
}

func TestNonExistingBBoxLabelShouldFail(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{MissingLabel: true}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{}, BoundingBoxes: []BoundingBoxRequest{{Label: "a-label"}}}, p)
	if !p.GotLabelNotFoundErr || p.GotSuccess {
		t.Fatal("expected label not found error, but go none")
	}
}

func TestHandleBoundingBoxValidationError(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: -2, Height: -4}}}, p)
	if !p.GotValidationErr || p.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestHandleAddBoundingBoxInternalErr(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{ErrOnAddBoundingBox: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: 2, Height: 4}}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAddBoundingBoxToImage(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, &i.FakeArtefactRepo{})
	itr.Execute(Request{Reader: &i.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: 2, Height: 4}}}, p)
	if !p.GotSuccess || repo.NumBoundingboxesAdded != 1 {
		t.Fatalf("expected to add one bounding box to repo, but got %v", repo.NumBoundingboxesAdded)
	}
}

func TestInternalErrOnAddLabelMustDeleteImage(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{ErrOnAddLabel: true, Err: e.ErrInternal}
	artefactRepo := &i.FakeArtefactRepo{}
	itr := NewInteractor(repo, artefactRepo)
	itr.Execute(Request{Reader: &i.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if repo.NumDeletedImages != 1 || artefactRepo.NumDeletedImages != 1 {
		t.Fatal("expected to delete image")
	}
}
