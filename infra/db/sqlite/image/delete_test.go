package sqlite

import (
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestHandleInternalErrOnDeleteImage(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	repo.Db.Close()
	err := repo.Delete(im.NewImageId())
	if err != e.ErrInternal {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestDeleteImage(t *testing.T) {
	repo := NewTestSQLiteImageRepo()
	id := im.NewImageId()
	repo.AddImage(id, "")
	err := repo.Delete(id)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
