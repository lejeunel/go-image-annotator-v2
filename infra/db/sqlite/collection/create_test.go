package sqlite

import (
	"errors"
	"testing"

	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func createCollection(repo *SQLiteCollectionRepo, name string) (*clc.Collection, error) {
	c := clc.NewCollection(clc.NewCollectionId(), name, clc.WithDescription("a-description"))
	if err := repo.Create(*c); err != nil {
		return nil, err
	}
	return c, nil

}

func TestInternalErrOnCreateShouldFail(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	repo.Db.Close()
	_, err := createCollection(repo, "a-collection")
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestCreate(t *testing.T) {
	_, err := createCollection(NewTestSQLiteCollectionRepo(), "a-collection")
	if err != nil {
		t.Fatalf("expected no error on create but got %v", err)
	}

}
