package scroller

import (
	"errors"
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func TestErrOnCurrentImageShouldFail(t *testing.T) {
	_, err := NewScroller(&FakeRepo{ErrOnImageExists: true, Err: e.ErrNotFound},
		im.NewImageId())
	if !errors.Is(err, e.ErrNotFound) {
		t.Fatalf("expected not found error, got %v", err)
	}
}

func TestNonExistingCollectionShouldFail(t *testing.T) {
	_, err := NewScroller(&FakeRepo{ErrOnCollectionExists: true, Err: e.ErrNotFound},
		im.NewImageId(), WithCollection("non-existing-collection"))
	if !errors.Is(err, e.ErrNotFound) {
		t.Fatalf("expected not found error, got %v", err)
	}
}

func TestSingleImageHasNoNextImage(t *testing.T) {
	s, _ := NewScroller(&FakeRepo{}, im.NewImageId())
	state, _ := s.State()
	if state.Next != nil {
		t.Fatal("expected no next image")
	}
}

func TestSingleImageHasNoPreviousImage(t *testing.T) {
	s, _ := NewScroller(&FakeRepo{}, im.NewImageId())
	state, _ := s.State()
	if state.Previous != nil {
		t.Fatal("expected no next image")
	}
}

func TestNextImage(t *testing.T) {
	next := &im.BaseImage{ImageId: im.NewImageId()}
	s, _ := NewScroller(&FakeRepo{NextImage: next}, im.NewImageId())
	state, _ := s.State()
	if state.Next == nil {
		t.Fatal("expected to get one next image")
	}
	if state.Next.ImageId != next.ImageId {
		t.Fatalf("expected to get next image with id %v, got %v", next.ImageId, state.Next.ImageId)
	}
}

func TestPreviousImage(t *testing.T) {
	prev := &im.BaseImage{ImageId: im.NewImageId()}
	s, _ := NewScroller(&FakeRepo{PreviousImage: prev}, im.NewImageId())
	state, _ := s.State()
	if state.Previous == nil {
		t.Fatal("expected to get one previous image")
	}
	if state.Previous.ImageId != prev.ImageId {
		t.Fatalf("expected to get previous image with id %v, got %v", prev.ImageId, state.Previous.ImageId)
	}
}

func TestStateContainsCriteria(t *testing.T) {
	next := &im.BaseImage{ImageId: im.NewImageId()}
	collection := "a-collection"
	s, _ := NewScroller(&FakeRepo{NextImage: next},
		im.NewImageId(), WithCollection(collection))
	state, _ := s.State()
	if state.Next.Collection != collection {
		t.Fatalf("expected next image to be in collection %v, got %v",
			collection, state.Next.Collection)
	}
}
