package sqlite

import (
	"fmt"

	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	s "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/update"
)

type SQLiteCollectionRepo struct {
	Db *sqlx.DB
}

type Row struct {
	Id          clc.CollectionId `db:"id"`
	Name        string           `db:"name"`
	Description string           `db:"description"`
}

func (r *SQLiteCollectionRepo) Create(c clc.Collection) error {
	query := "INSERT INTO collections (id, name, description) VALUES ($1,$2,$3)"
	_, err := r.Db.Exec(query, c.Id.String(), c.Name, c.Description)
	if err != nil {
		return fmt.Errorf("%v: %w", err, e.ErrInternal)
	}

	return nil
}
func (r *SQLiteCollectionRepo) FindCollectionByName(name string) (*clc.Collection, error) {

	row := Row{}
	err := r.Db.Get(&row,
		"SELECT id,name,description FROM collections WHERE name=$1", name)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, e.ErrNotFound
		default:
			return nil, fmt.Errorf("%v: %w", err, e.ErrInternal)
		}
	}

	return clc.NewCollection(row.Id, row.Name, clc.WithDescription(row.Description)), nil
}

func (r *SQLiteCollectionRepo) Exists(name string) (bool, error) {
	var exists bool

	err := r.Db.Get(&exists, `SELECT EXISTS (SELECT 1 FROM collections WHERE name = $1)`, name)
	if err != nil {
		return false, e.ErrInternal
	}

	return exists, nil
}

func (r *SQLiteCollectionRepo) Delete(name string) error {
	_, err := r.Db.Exec("DELETE FROM collections WHERE name=$1", name)

	if err != nil {
		return fmt.Errorf("%v: %w", err, e.ErrInternal)
	}
	return nil
}

func (r *SQLiteCollectionRepo) Update(m update.Model) error {
	query := "UPDATE collections SET name=$1,description=$2 WHERE name=$3"
	_, err := r.Db.Exec(query, m.NewName, m.NewDescription, m.Name)

	if err != nil {
		return fmt.Errorf("%v: %w", err, e.ErrInternal)
	}

	return nil
}
func (r *SQLiteCollectionRepo) IsPopulated(name string) (*bool, error) {
	var count int64

	var query string
	var err error
	query = "SELECT COUNT(*) FROM images_collections WHERE collection_id=(SELECT id FROM collections WHERE name=$1)"
	err = r.Db.QueryRow(query, name).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", err, e.ErrInternal)
	}
	isPopulated := count > 0
	return &isPopulated, nil
}

func NewSQLiteCollectionRepo(db *sqlx.DB) *SQLiteCollectionRepo {
	return &SQLiteCollectionRepo{Db: db}
}

func NewTestSQLiteCollectionRepo() *SQLiteCollectionRepo {
	return NewSQLiteCollectionRepo(s.NewSQLiteDB(":memory:"))
}
