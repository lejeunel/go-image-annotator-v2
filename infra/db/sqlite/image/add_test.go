package sqlite

import (
	"errors"
	"testing"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	s "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	cr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
)

func TestInternalErrOnCreateShouldFail(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	repo.Db.Close()
	err := repo.AddImageToCollection(im.NewImageId(), clc.NewCollectionId())
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestInternalErrOnImageIsInCollectionShouldFail(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	repo.Db.Close()
	_, err := repo.ImageExistsInCollection(im.NewImageId(), clc.NewCollectionId())
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestAddImageInCollection(t *testing.T) {
	db := s.NewSQLiteDB(":memory:")
	collectionRepo := cr.NewSQLiteCollectionRepo(db)
	imageRepo := NewSQLiteImageRepo(db)

	collectionId := clc.NewCollectionId()
	collectionRepo.Create(*clc.NewCollection(collectionId, "a-collection"))

	imageId := im.NewImageId()
	err := imageRepo.AddImageToCollection(imageId, collectionId)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	isAdded, err := imageRepo.ImageExistsInCollection(imageId, collectionId)
	if err != nil {
		t.Fatalf("expected no error when checking existence of image in collection, got %v", err)
	}
	if !isAdded {
		t.Fatal("expected that image is added to collection")
	}
}
