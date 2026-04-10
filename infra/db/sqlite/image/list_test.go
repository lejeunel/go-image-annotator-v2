package sqlite

import (
	"errors"
	ist "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	s "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	sc "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	"testing"
)

type ImageListingTestingRepos struct {
	Image      SQLiteImageRepo
	Collection sc.SQLiteCollectionRepo
}

func CreateSingleImageCollection(repos ImageListingTestingRepos, collectionName string) (im.Image, clc.Collection) {
	collection := clc.NewCollection(clc.NewCollectionId(), collectionName)
	repos.Collection.Create(*collection)
	imageId := im.NewImageId()
	image := im.NewImage(imageId, *collection)
	repos.Image.AddImage(image.Id, "the-hash", "the-mimetype")
	repos.Image.AddImageToCollection(image.Id, collection.Id)
	return *image, *collection
}

func NewImageListingTestRepos() ImageListingTestingRepos {
	db := s.NewSQLiteDB(":memory:")
	return ImageListingTestingRepos{Image: *NewSQLiteImageRepo(db),
		Collection: *sc.NewSQLiteCollectionRepo(db),
	}
}

func TestInternalErrOnImageListShouldFail(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	repo.Db.Close()
	_, err := repo.List(ist.FilteringParams{})
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestListOneImage(t *testing.T) {
	repos := NewImageListingTestRepos()
	collectionName := "a-collection"
	collection := clc.NewCollection(clc.NewCollectionId(), collectionName)
	repos.Collection.Create(*collection)
	image := im.NewImage(im.NewImageId(), *collection)
	repos.Image.AddImage(image.Id, "", "")
	repos.Image.AddImageToCollection(image.Id, collection.Id)

	r, _ := repos.Image.List(ist.FilteringParams{PageSize: 2, Page: 1})
	if len(*r) != 1 {
		t.Fatalf("expected to retrieve one image, got %v", len(*r))
	}
}

func TestListOneImageInGivenCollection(t *testing.T) {
	repos := NewImageListingTestRepos()

	firstImage, firstCollection := CreateSingleImageCollection(repos, "first-collection")
	CreateSingleImageCollection(repos, "second-collection")

	r, _ := repos.Image.List(ist.FilteringParams{Collection: &firstCollection.Name, PageSize: 2, Page: 1})
	if len(*r) != 1 {
		t.Fatalf("expected to retrieve one image, got %v", len(*r))
	}
	images := *r
	if images[0].ImageId != firstImage.Id {
		t.Fatalf("expected to retrieve first image with id %v, got %v", firstImage.Id, images[0].ImageId)
	}
	if images[0].Collection != firstCollection.Name {
		t.Fatalf("expected that image belongs to collection named %v, got %v", firstCollection.Name, images[0].Collection)
	}
}
