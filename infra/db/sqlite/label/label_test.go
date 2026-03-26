package sqlite

import (
	"testing"

	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	s "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
)

func NewTestSQLiteLabelRepo() *SQLiteLabelRepo {
	return NewSQLiteLabelRepo(s.NewSQLiteDB(":memory:"))

}

func createLabel(repo *SQLiteLabelRepo) (*lbl.Label, error) {
	id := lbl.NewLabelID()
	label := lbl.NewLabel(id, "a-label", lbl.WithDescription("a-description"))
	if err := repo.Create(*label); err != nil {
		return nil, err
	}
	return label, nil

}

func TestCreate(t *testing.T) {
	repo := NewTestSQLiteLabelRepo()
	_, err := createLabel(repo)
	if err != nil {
		t.Fatalf("expected no error on create but got %v", err)
	}

}

func TestRetrieveNonExistingShouldFail(t *testing.T) {
	repo := NewTestSQLiteLabelRepo()
	createLabel(repo)
	_, err := repo.Find("non-existing-label")
	if err != e.ErrNotFound {
		t.Fatalf("expected not found error, got %v", err)
	}
}

func TestInternalErrOnCreateShouldFail(t *testing.T) {
	repo := NewTestSQLiteLabelRepo()
	repo.Db.Close()
	_, err := createLabel(repo)
	if err != e.ErrInternal {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestInternalErrOnFindShouldFail(t *testing.T) {
	repo := NewTestSQLiteLabelRepo()
	createLabel(repo)
	repo.Db.Close()
	_, err := repo.Find("a-label")
	if err != e.ErrInternal {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestRetrieve(t *testing.T) {
	repo := NewTestSQLiteLabelRepo()
	label, _ := createLabel(repo)
	r, err := repo.Find("a-label")
	if err != nil {
		t.Fatalf("expected no error on find, got %v", err)
	}
	if (r.Name != label.Name) || (r.Description != label.Description) || (r.Id != label.Id) {
		t.Fatalf("expected to retrieve name %v, description %v, and id %v, got %v, %v, %v",
			label.Name, label.Description, label.Id, r.Name, r.Description, r.Id)

	}

}
