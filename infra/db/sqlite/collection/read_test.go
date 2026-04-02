package sqlite

import (
	"errors"
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func TestRetrieveNonExistingShouldFail(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	createCollection(repo, "a-collection")
	_, err := repo.FindCollectionByName("non-existing-collection")
	if !errors.Is(err, e.ErrNotFound) {
		t.Fatalf("expected not found error, got %v", err)
	}
}

func TestInternalErrOnFindShouldFail(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	createCollection(repo, "a-collection")
	repo.Db.Close()
	_, err := repo.FindCollectionByName("a-collection")
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestRetrieve(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	collection, _ := createCollection(repo, "a-collection")
	r, err := repo.FindCollectionByName("a-collection")
	if err != nil {
		t.Fatalf("expected no error on find, got %v", err)
	}
	if (r.Name != collection.Name) || (r.Description != collection.Description) || (r.Id != collection.Id) {
		t.Fatalf("expected to retrieve name %v, description %v, and id %v, got %v, %v, %v",
			collection.Name, collection.Description, collection.Id, r.Name, r.Description, r.Id)

	}

}
