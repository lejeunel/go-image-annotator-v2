package sqlite

import (
	"errors"
	ist "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestCountAddedImageToCollection(t *testing.T) {
	repos := NewImageTestRepos()
	collection := "a-collection"
	AddImageToCollection(repos, collection, "")
	count, err := repos.Image.Count(ist.CountingParams{Collection: &collection})
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
	count, err := repos.Image.Count(ist.CountingParams{})
	if err != nil {
		t.Fatalf("expected no error when counting images in collection, got %v", err)
	}
	if *count != 1 {
		t.Fatalf("expected that one image is added to collection, got %v", *count)
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
