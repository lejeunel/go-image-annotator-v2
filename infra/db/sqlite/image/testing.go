package sqlite

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	s "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	cr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
)

type ImageTestingRepos struct {
	Image      SQLiteImageRepo
	Collection cr.SQLiteCollectionRepo
}

func NewImageTestRepos() ImageTestingRepos {
	db := s.NewSQLiteDB(":memory:")
	return ImageTestingRepos{Image: *NewSQLiteImageRepo(db),
		Collection: *cr.NewSQLiteCollectionRepo(db)}
}

func AddImageToCollection(repos ImageTestingRepos, collectionName string, hash string) (*im.ImageId, *clc.CollectionId, error) {
	collectionId := clc.NewCollectionId()
	repos.Collection.Create(*clc.NewCollection(collectionId, collectionName))
	imageId := im.NewImageId()
	repos.Image.AddImage(imageId, "the-hash", "the-mimetype")

	return &imageId, &collectionId, repos.Image.AddImageToCollection(imageId, collectionId)
}
