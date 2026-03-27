package sqlite

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	u "github.com/lejeunel/go-image-annotator-v2/use-cases/label/update"
	"testing"
)

func TestInternalErrOnLabelUpdateShouldFail(t *testing.T) {
	repo := NewTestSQLiteLabelRepo()
	repo.Db.Close()
	err := repo.Update(u.Model{})
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestUpdateLabel(t *testing.T) {
	repo := NewTestSQLiteLabelRepo()
	label, _ := createLabel(repo, "a-label")
	err := repo.Update(u.Model{Name: label.Name, NewName: "new-label-name", NewDescription: "new-description"})
	if err != nil {
		t.Fatalf("did not expect error, got %v", err)
	}
}
