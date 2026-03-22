package labeling

import (
	"errors"
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingLabelShouldFail(t *testing.T) {
	s := NewLabelingService(&FakeRepo{ErrOnFindLabel: true, Err: e.ErrNotFound})
	_, err := s.Init(im.NewImageID(), "a-collection", "a-label")
	if !errors.Is(err, e.ErrNotFound) {
		t.Fatalf("expected error not found, got %v", err)
	}
}

func TestInternalErrOnFindLabelShouldFail(t *testing.T) {
	s := NewLabelingService(&FakeRepo{ErrOnFindLabel: true, Err: e.ErrInternal})
	_, err := s.Init(im.NewImageID(), "a-collection", "a-label")
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestNonExistingCollectionShouldFail(t *testing.T) {
	s := NewLabelingService(&FakeRepo{ErrOnFindCollection: true, Err: e.ErrNotFound})
	_, err := s.Init(im.NewImageID(), "a-collection", "a-label")
	if !errors.Is(err, e.ErrNotFound) {
		t.Fatalf("expected error not found, got %v", err)
	}
}

func TestInternalErrOnFindCollectionShouldFail(t *testing.T) {
	s := NewLabelingService(&FakeRepo{ErrOnFindCollection: true, Err: e.ErrInternal})
	_, err := s.Init(im.NewImageID(), "a-collection", "a-label")
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected error not found, got %v", err)
	}
}

func TestImageNotInCollectionShouldFail(t *testing.T) {
	s := NewLabelingService(&FakeRepo{ImageNotInCollection: true})
	_, err := s.Init(im.NewImageID(), "a-collection", "a-label")
	if !errors.Is(err, e.ErrDependency) {
		t.Fatalf("expected dependency error, got %v", err)
	}
}

func TestInternalErrOnImageIsInCollectionShouldFail(t *testing.T) {
	s := NewLabelingService(&FakeRepo{ErrOnImageIsInCollection: true, Err: e.ErrInternal})
	_, err := s.Init(im.NewImageID(), "a-collection", "a-label")
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestSuccess(t *testing.T) {
	s := NewLabelingService(&FakeRepo{})
	ctx, err := s.Init(im.NewImageID(), "a-collection", "a-label")
	if err != nil {
		t.Fatalf("expected no error")
	}
	if ctx == nil {
		t.Fatal("expected a labeling context, but got nil")
	}
}
