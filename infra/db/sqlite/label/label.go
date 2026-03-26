package sqlite

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type SQLiteLabelRepo struct {
	Db *sqlx.DB
}

type LabelRecord struct {
	Id          lbl.LabelId `db:"id"`
	Name        string      `db:"name"`
	Description string      `db:"description"`
}

func (r *SQLiteLabelRepo) Create(l lbl.Label) error {
	query := "INSERT INTO labels (id, name, description) VALUES ($1,$2,$3)"
	_, err := r.Db.Exec(query, l.Id.String(), l.Name, l.Description)
	if err != nil {
		return e.ErrInternal
	}

	return nil

}

func (r *SQLiteLabelRepo) Find(name string) (*lbl.Label, error) {
	record := LabelRecord{}
	err := r.Db.Get(&record,
		"SELECT id,name,description FROM labels WHERE name=$1", name)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, e.ErrNotFound
		default:
			return nil, e.ErrInternal
		}

	}

	return lbl.NewLabel(record.Id, record.Name, lbl.WithDescription(record.Description)), nil
}

func NewSQLiteLabelRepo(db *sqlx.DB) *SQLiteLabelRepo {
	return &SQLiteLabelRepo{Db: db}
}
