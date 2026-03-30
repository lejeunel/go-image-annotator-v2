package sqlite

import (
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestRetrieveImageIdByHash(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	imageId := im.NewImageId()
	hash := "the-hash"
	err := repo.AddImage(imageId, hash)
	if err != nil {
		t.Fatalf("expected no error on adding image, got %v", err)
	}

	existingId, err := repo.FindImageIdByHash(hash)
	if err != nil {
		t.Fatalf("expected no error finding image by hash, got %v", err)
	}
	if *existingId != imageId {
		t.Fatalf("expected to retrieve image with identical hash and id %v, got %v", imageId, existingId)
	}
}

func TestRetrieveImageIdByNonExistingHashShouldFail(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	imageId := im.NewImageId()
	hash := "the-hash"
	repo.AddImage(imageId, hash)
	_, err := repo.FindImageIdByHash("non-existing-hash")
	if err != e.ErrNotFound {
		t.Fatalf("expected not found error, got %v", err)
	}
}

func TestRetrieveImageIdByInternalErrShouldFail(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	repo.Db.Close()
	_, err := repo.FindImageIdByHash("")
	if err != e.ErrInternal {
		t.Fatalf("expected internal error, got %v", err)
	}
}
