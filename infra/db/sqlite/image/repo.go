package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	ist "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
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
		return fmt.Errorf("inserting record into image to collection junction table: %v: %w", err, e.ErrInternal)
	}

	return nil
}
func (r *SQLiteImageRepo) Count(f ist.CountingParams) (*int64, error) {
	var count int64

	var query string
	var err error
	if f.Collection != nil {
		query = "SELECT COUNT(*) FROM images_collections WHERE collection_id=(SELECT id FROM collections WHERE name=$1)"
		err = r.Db.QueryRow(query, f.Collection).Scan(&count)
	} else {
		query = "SELECT COUNT(*) FROM images"
		err = r.Db.QueryRow(query).Scan(&count)
	}
	if err != nil {
		return nil, fmt.Errorf("counting image records: %v: %w", err, e.ErrInternal)
	}
	return &count, nil

}
func (r *SQLiteImageRepo) ImageExistsInCollection(imageId im.ImageId, collectionId clc.CollectionId) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM images_collections WHERE image_id=$1 AND collection_id=$2"
	err := r.Db.QueryRow(query, imageId.String(), collectionId.String()).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("checking image to collection junction records: %v: %w", err, e.ErrInternal)
	}

	return count > 0, nil
}
func (r *SQLiteImageRepo) ImageExists(imageId im.ImageId) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM images WHERE id=$1"
	err := r.Db.QueryRow(query, imageId.String()).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("checking that image exists: %v: %w", err, e.ErrInternal)
	}

	return count > 0, nil
}

func (r *SQLiteImageRepo) AddImage(imageId im.ImageId, hash string) error {
	query := "INSERT INTO images (id, hash) VALUES ($1,$2)"
	_, err := r.Db.Exec(query, imageId.String(), hash)
	if err != nil {
		return fmt.Errorf("inserting image record: %v: %w", err, e.ErrInternal)
	}
	return nil
}
func (r *SQLiteImageRepo) FindImageIdByHash(hash string) (*im.ImageId, error) {
	errCtx := "finding image record by hash"
	var imageId im.ImageId
	err := r.Db.Get(&imageId, "SELECT id FROM images WHERE hash = $1", hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%v: %v: %w", errCtx, err, e.ErrNotFound)
		}
		return nil, fmt.Errorf("%v: %v: %w", errCtx, err, e.ErrInternal)
	}
	return &imageId, nil
}
func (r *SQLiteImageRepo) Delete(id im.ImageId) error {
	_, err := r.Db.Exec("DELETE FROM images WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("deleting image record: %v: %w", err, e.ErrInternal)
	}
	return nil
}
func (r *SQLiteImageRepo) RemoveImageFromCollection(imageId im.ImageId, collectionId clc.CollectionId) error {
	_, err := r.Db.Exec("DELETE FROM images_collections WHERE image_id = $1 AND collection_id = $2",
		imageId, collectionId)
	if err != nil {
		return fmt.Errorf("removing image from image to collection junction table: %v: %w", err, e.ErrInternal)
	}
	return nil
}

func NewSQLiteImageRepo(db *sqlx.DB) *SQLiteImageRepo {
	return &SQLiteImageRepo{Db: db}
}

func NewTestSQLiteImageRepo() *SQLiteImageRepo {
	return NewSQLiteImageRepo(s.NewSQLiteDB(":memory:"))
}
