package unassign_label

import (
	lbi "github.com/lejeunel/go-image-annotator-v2/domain/labeling"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestUnassignNonExistingLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{ErrOnFindLabel: true, Err: e.ErrNotFound})
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected label not found error")
	}
}
