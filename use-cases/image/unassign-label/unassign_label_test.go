package unassign_label

import (
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbi "github.com/lejeunel/go-image-annotator-v2/domain/labeling"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestHandleNotFoundErr(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := &lbi.FakeLabelingService{Err: e.ErrNotFound}
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected not found error")
	}
}

func TestHandleDependencyErr(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := &lbi.FakeLabelingService{Err: e.ErrDependency}
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotDependencyErr || presenter.GotSuccess {
		t.Fatal("expected dependency error")
	}
}
func TestHandleInternalErr(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := &lbi.FakeLabelingService{Err: e.ErrInternal}
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestUnassignLabelToImage(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{})
	itr := NewInteractor(repo, presenter, labelingService)
	itr.Execute(Request{Label: "a-label", ImageId: im.NewImageID(), Collection: "a-collection"})
	if !presenter.GotSuccess || !repo.RemovedLabel {
		t.Fatal("expected to remove label")
	}
}
