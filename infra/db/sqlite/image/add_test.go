package sqlite

import (
	"errors"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestInternalErrOnImageIsInCollectionShouldFail(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	repo.Db.Close()
	_, err := repo.ImageExistsInCollection(im.NewImageId(), clc.NewCollectionId())
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestCountAddedImageToCollection(t *testing.T) {
	repos := NewImageTestRepos()
	collection := "a-collection"
	AddImageToCollection(repos, collection, "")
	count, err := repos.Image.Count(im.CountingParams{Collection: &collection})
	if err != nil {
		t.Fatalf("expected no error when counting images in collection, got %v", err)
	}
	if *count != 1 {
		t.Fatalf("expected that one image is added to collection, got %v", *count)
	}
}

func TestCountAllImagesWhenAddingImageToCollection(t *testing.T) {
	repos := NewImageTestRepos()
	AddImageToCollection(repos, "a-collection", "")
	count, err := repos.Image.Count(im.CountingParams{})
	if err != nil {
		t.Fatalf("expected no error when counting images in collection, got %v", err)
	}
	if *count != 1 {
		t.Fatalf("expected that one image is added to collection, got %v", *count)
	}
}

func TestAddedImageToCollectionExists(t *testing.T) {
	repos := NewImageTestRepos()
	imageId, collectionId, _ := AddImageToCollection(repos, "a-collectxion", "the-hash")
	isAdded, err := repos.Image.ImageExistsInCollection(*imageId, *collectionId)
	if err != nil {
		t.Fatalf("expected no error when checking existence of image in collection, got %v", err)
	}
	if !isAdded {
		t.Fatal("expected that image is added to collection")
	}
}

func TestInternalErrOnCreateShouldFail(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	repo.Db.Close()
	err := repo.AddImageToCollection(im.NewImageId(), clc.NewCollectionId())
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestInternalErrOnIsCollectionPopulatedShouldFail(t *testing.T) {
	repos := NewImageTestRepos()
	collectionName := "a-collection"
	AddImageToCollection(repos, collectionName, "the-hash")
	repos.Image.Db.Close()
	_, err := repos.Collection.IsPopulated(collectionName)
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestIsCollectionPopulated(t *testing.T) {
	repos := NewImageTestRepos()
	collectionName := "a-collection"
	AddImageToCollection(repos, collectionName, "the-hash")
	isPopulated, err := repos.Collection.IsPopulated(collectionName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !(*isPopulated) {
		t.Fatal("expected populated collection, got")
	}
}
