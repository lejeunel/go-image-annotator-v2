package assign_label

import (
	"testing"

	lbi "github.com/lejeunel/go-image-annotator-v2/domain/labeling"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
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

func TestAssignLabelToImage(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	labelingService := &lbi.FakeLabelingService{}
	itr := NewInteractor(repo, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotSuccess || !repo.GotLabel {
		t.Fatal("expected to add one label")
	}
}
