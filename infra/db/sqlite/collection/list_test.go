package sqlite

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	l "github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	"testing"
)

func TestInternalErrOnLabelCountShouldFail(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	repo.Db.Close()
	_, err := repo.Count()
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestCountLabels(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	createCollection(repo, "a-collection")
	count, _ := repo.Count()
	if *count != 1 {
		t.Fatalf("expected label count %v, got %v", 1, count)
	}
}

func TestInternalErrOnLabelListShouldFail(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	repo.Db.Close()
	_, err := repo.List(l.Request{})
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestListCollections(t *testing.T) {
	repo := NewTestSQLiteCollectionRepo()
	firstCollection, _ := createCollection(repo, "a-collection")
	secondCollection, _ := createCollection(repo, "another-collection")
	collections, err := repo.List(l.Request{Page: 1, PageSize: 2})
	if err != nil {
		t.Fatalf("did not expect error, got %v", err)
	}
	if len(collections) != 2 {
		t.Fatalf("expected two collections, got %v", len(collections))
	}
	if collections[0].Name == collections[1].Name {
		t.Fatalf("expected to retrieve two distinct collections with name %v and %v, got %v and %v",
			firstCollection.Name, secondCollection.Name, collections[0].Name, collections[1].Name)

	}
	if collections[0].Description != firstCollection.Description {
		t.Fatalf("expected to retrieve collection with description %v , got %v",
			firstCollection.Description, collections[0].Description)
	}
}
