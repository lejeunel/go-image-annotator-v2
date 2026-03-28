package sqlite

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	s "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
)

type SQLiteImageRepo struct {
	Db *sqlx.DB
}

type Row struct {
	Id im.ImageId `db:"id"`
}

func (r *SQLiteImageRepo) AddImageToCollection(imageId im.ImageId, collectionId clc.CollectionId) error {
	query := "INSERT INTO images_collections (image_id, collection_id) VALUES ($1,$2)"
	_, err := r.Db.Exec(query, imageId.String(), collectionId.String())
	if err != nil {
		return fmt.Errorf("%v: %w", err, e.ErrInternal)
	}

	return nil
}
func (r *SQLiteImageRepo) ImageExistsInCollection(imageId im.ImageId, collectionId clc.CollectionId) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM images_collections WHERE image_id=$1 AND collection_id=$2"
	err := r.Db.QueryRow(query, imageId.String(), collectionId.String()).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("%v: %w", err, e.ErrInternal)
	}

	return count > 0, nil
}
func NewSQLiteImageRepo(db *sqlx.DB) *SQLiteImageRepo {
	return &SQLiteImageRepo{Db: db}
}

func NewTestSQLiteImageRepo() *SQLiteImageRepo {
	return NewSQLiteImageRepo(s.NewSQLiteDB(":memory:"))
}
