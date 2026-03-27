package sqlite

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	u "github.com/lejeunel/go-image-annotator-v2/use-cases/collection/update"
	"testing"
)

func TestInternalErrOnCollectionUpdateShouldFail(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	repo.Db.Close()
	err := repo.Update(u.Model{})
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestUpdateCollection(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	collection, _ := createCollection(repo, "a-collection")
	err := repo.Update(u.Model{Name: collection.Name, NewName: "new-collection-name", NewDescription: "new-description"})
	if err != nil {
		t.Fatalf("did not expect error, got %v", err)
	}
}
