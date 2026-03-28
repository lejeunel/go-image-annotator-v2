package image

import (
	"bytes"
	"errors"
	"io"
	"testing"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingCollectionShouldFail(t *testing.T) {
	s := NewImageStore(&FakeImageRepo{}, &FakeCollectionRepo{MissingCollection: true},
		&FakeAnnotationRepo{}, &FakeArtefactRepo{})
	_, err := s.Find(BaseImage{NewImageId(), "a-collection"})
	if !errors.Is(err, e.ErrNotFound) {
		t.Fatalf("expected error not found, got %v", err)
	}
}

func TestErrOnFindLabelShouldFail(t *testing.T) {
	s := NewImageStore(&FakeImageRepo{}, &FakeCollectionRepo{},
		&FakeAnnotationRepo{ErrOnFindImageLabel: true, Err: e.ErrInternal}, &FakeArtefactRepo{})
	_, err := s.Find(BaseImage{NewImageId(), "a-collection"})
	if err == nil {
		t.Fatalf("expected error, got %v", err)
	}
}

func TestErrOnFindBoundingBoxesShouldFail(t *testing.T) {
	s := NewImageStore(&FakeImageRepo{}, &FakeCollectionRepo{},
		&FakeAnnotationRepo{ErrOnFindBoundingBoxes: true, Err: e.ErrInternal}, &FakeArtefactRepo{})
	_, err := s.Find(BaseImage{NewImageId(), "a-collection"})
	if err == nil {
		t.Fatalf("expected error, got %v", err)
	}
}

func TestErrOnExistsShouldFail(t *testing.T) {
	s := NewImageStore(&FakeImageRepo{ErrOnExists: true, Err: e.ErrInternal}, &FakeCollectionRepo{},
		&FakeAnnotationRepo{ErrOnFindBoundingBoxes: true, Err: e.ErrInternal}, &FakeArtefactRepo{})
	_, err := s.Find(BaseImage{NewImageId(), "a-collection"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestFindImage(t *testing.T) {
	label := lbl.NewLabel(lbl.NewLabelId(), "a-label")
	labels := []*a.ImageLabel{{Id: a.NewAnnotationId(), Label: *label}}
	bboxes := []*a.BoundingBox{{Id: a.NewAnnotationId(), Label: *label}}
	collection := clc.NewCollection(clc.NewCollectionId(), "a-collection")
	data := []byte("test-data")

	s := NewImageStore(&FakeImageRepo{}, &FakeCollectionRepo{Collection: *collection},
		&FakeAnnotationRepo{Labels: labels, BoundingBoxes: bboxes}, &FakeArtefactRepo{Data: data})
	image, _ := s.Find(BaseImage{ImageId: NewImageId(), Collection: collection.Name})
	if !(image.Collection.Id == collection.Id) {
		t.Fatalf("expected to retrieve image in collection %v, got %v ",
			collection.Id, image.Collection.Id)
	}
	if !(len(image.Labels) == 1) {
		t.Fatalf("expected to retrieve image with 1 label, got %v", len(image.Labels))
	}
	if !(len(image.BoundingBoxes) == 1) {
		t.Fatalf("expected to retrieve image with 1 bounding box, got %v", len(image.BoundingBoxes))
	}
	gotBytes, _ := io.ReadAll(image.Reader)
	if !bytes.Equal(gotBytes, data) {
		t.Fatalf("expected to retrieve bytes %v, got %v", data, gotBytes)

	}
}
