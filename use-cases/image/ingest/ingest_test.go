package ingest

import (
	"testing"

	ast "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingCollectionShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{}, &FakeCollectionRepo{MissingCollection: true}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{}, p)
	if !p.GotCollectionNotFoundErr || p.GotSuccess {
		t.Fatal("expected not found error")
	}
}

func TestInvalidImageDataShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{}, &FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{Err: e.ErrInternal}}, p)
	if !p.GotInvalidImageDataErr || p.GotSuccess {
		t.Fatal("expected invalid data error")
	}
}

func TestHandleInternalErrorOnCollectionExistsCheck(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{ErrOnFindCollection: true, Err: e.ErrInternal}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleArtefactRepoError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{ErrOnFindCollection: true, Err: e.ErrInternal}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{Err: e.ErrInternal}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected invalid data error")
	}
}

func TestNonExistingLabelShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{MissingLabel: true},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if !p.GotLabelNotFoundErr || p.GotSuccess {
		t.Fatal("expected label not found error, but go none")
	}
}

func TestHandleLabelExistsInternalErr(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{ErrOnLabelExists: true, Err: e.ErrInternal},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{Err: e.ErrInternal}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleIngestionInternalErr(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{ErrOnAddImageToCollection: true, Err: e.ErrInternal},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{Err: e.ErrInternal}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleAddLabelInternalErr(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{ErrOnAddLabel: true, Err: e.ErrInternal}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestHandleValidationErrorOnAddLabel(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}, Labels: []string{"a-label", "a-label"}}, p)
	if !p.GotValidationErr || p.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestAddImageDuplicateHashShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{HashAlreadyExists: true},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}}, p)
	if !p.GotDuplicateImage || p.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestHandleDuplicateHashInternalErr(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{ErrOnFindHash: true, Err: e.ErrInternal},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestNonExistingBBoxLabelShouldFail(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{MissingLabel: true},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}, BoundingBoxes: []BoundingBoxRequest{{Label: "a-label"}}}, p)
	if !p.GotLabelNotFoundErr || p.GotSuccess {
		t.Fatal("expected label not found error, but go none")
	}
}

func TestHandleBoundingBoxValidationError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: -2, Height: -4}}}, p)
	if !p.GotValidationErr || p.GotSuccess {
		t.Fatal("expected validation error")
	}
}

func TestHandleAddBoundingBoxInternalErr(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{ErrOnAddBoundingBox: true, Err: e.ErrInternal}, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: 2, Height: 4}}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestInternalErrOnAddLabelMustDeleteImage(t *testing.T) {
	p := &FakePresenter{}
	artefactRepo := &ast.FakeArtefactRepo{}
	imageRepo := &FakeImageRepo{}
	itr := NewInteractor(imageRepo,
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{ErrOnAddLabel: true, Err: e.ErrInternal}, artefactRepo, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if imageRepo.NumDeletedImages != 1 || artefactRepo.NumDeletedImages != 1 {
		t.Fatal("expected to delete image")
	}
}

func TestAddBoundingBoxToImage(t *testing.T) {
	p := &FakePresenter{}
	annotationRepo := &FakeAnnotationRepo{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		annotationRepo, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{},
		BoundingBoxes: []BoundingBoxRequest{{Label: "a-label", Xc: 10, Yc: 10, Width: 2, Height: 4}}}, p)
	if !p.GotSuccess || annotationRepo.NumBoundingboxesAdded != 1 {
		t.Fatalf("expected to add one bounding box to repo, but got %v", annotationRepo.NumBoundingboxesAdded)
	}
}

func TestInternalErrOnAddImageShouldFail(t *testing.T) {
	p := &FakePresenter{}
	artefactRepo := &ast.FakeArtefactRepo{}
	itr := NewInteractor(&FakeImageRepo{ErrOnAddImage: true, Err: e.ErrInternal},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		&FakeAnnotationRepo{}, artefactRepo, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}, Labels: []string{"a-label"}}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAddImageWithHash(t *testing.T) {
	p := &FakePresenter{}
	annotationRepo := &FakeAnnotationRepo{}
	hash := "the-hash"
	imageRepo := &FakeImageRepo{}
	itr := NewInteractor(imageRepo,
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		annotationRepo, &ast.FakeArtefactRepo{}, &FakeHasher{Hash_: hash})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}}, p)
	if !p.GotSuccess || imageRepo.GotHash != hash {
		t.Fatalf("expected to store image hash %v, got %v", hash, imageRepo.GotHash)
	}
}

func TestAddImageWithTwoLabels(t *testing.T) {
	p := &FakePresenter{}
	annotationRepo := &FakeAnnotationRepo{}
	itr := NewInteractor(&FakeImageRepo{},
		&FakeCollectionRepo{}, &FakeLabelRepo{},
		annotationRepo, &ast.FakeArtefactRepo{}, &FakeHasher{})
	itr.Execute(Request{Reader: &ast.FakeImageReader{}, Labels: []string{"a-label", "another-label"}}, p)
	if !p.GotSuccess || annotationRepo.NumLabelsAdded != 2 {
		t.Fatalf("expected to add two labels, but got %v", annotationRepo.NumLabelsAdded)
	}
}
