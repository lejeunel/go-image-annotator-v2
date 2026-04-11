package sqlite

import (
	"fmt"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	s "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	clcsql "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	imsql "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/image"
	"strings"
)

type SQLiteScrollerRepos struct {
	Scroller   *SQLiteScrollerRepo
	Image      *imsql.SQLiteImageRepo
	Collection *clcsql.SQLiteCollectionRepo
}

func NewTestScrollerRepos() *SQLiteScrollerRepos {
	db := s.NewSQLiteDB(":memory:")
	return &SQLiteScrollerRepos{Scroller: NewSQLiteScrollerRepo(db),
		Image:      imsql.NewSQLiteImageRepo(db),
		Collection: clcsql.NewSQLiteCollectionRepo(db)}
}

func FakeUUIDFromInt(n int) string {
	digit := fmt.Sprintf("%d", n)

	// Repeat the digit to fill 32 hex characters
	full := strings.Repeat(digit, 32)

	// Format as UUID: 8-4-4-4-12
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		full[0:8],
		full[8:12],
		full[12:16],
		full[16:20],
		full[20:32],
	)
}

func CreateImagesWithOrderedIds(repo *imsql.SQLiteImageRepo, num int) []im.ImageId {
	ids := []im.ImageId{}
	for n := range num {
		id, _ := im.NewImageIdFromString(FakeUUIDFromInt(n))
		repo.AddImage(id, id.String(), "")
		ids = append(ids, id)
	}
	return ids
}

func CreateImageInCollection(imRepo *imsql.SQLiteImageRepo, clcRepo *clcsql.SQLiteCollectionRepo,
	imageId im.ImageId, collectionName string) im.Image {
	collection := clc.NewCollection(clc.NewCollectionId(), collectionName)
	clcRepo.Create(*collection)
	image := im.NewImage(im.NewImageId(), *collection)
	imRepo.AddImage(image.Id, image.Id.String(), "")
	imRepo.AddImageToCollection(image.Id, image.Collection.Id)
	return *image
}
