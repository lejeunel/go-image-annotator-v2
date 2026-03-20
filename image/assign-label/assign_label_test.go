package assign_label

import (
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbi "github.com/lejeunel/go-image-annotator-v2/domain/labeling"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestAssignNonExistingLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{ErrOnFindLabel: true, Err: e.ErrNotFound})
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected label not found error")
	}
}

func TestNonExistingCollectionShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{ErrOnFindCollection: true, Err: e.ErrNotFound})
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatal("expected collection not found error")
	}
}

func TestInternalErrOnFindLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{ErrOnFindLabel: true, Err: e.ErrInternal})
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestInternalErrOnFindCollectionShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{ErrOnFindCollection: true, Err: e.ErrInternal})
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAssignLabelToImageNotInCollectionShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{ImageNotInCollection: true})
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{ImageId: im.NewImageID()})
	if !presenter.GotImageNotInCollectionErr || presenter.GotSuccess {
		t.Fatal("expected image not in collection error")
	}
}

func TestInternalErrOnImageIsInCollectionShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{ErrOnImageIsInCollection: true, Err: e.ErrInternal})
	itr := NewInteractor(&FakeRepo{}, presenter, labelingService)
	itr.Execute(Request{ImageId: im.NewImageID()})
	if !presenter.GotInternalErr || presenter.GotSuccess || presenter.GotImageNotInCollectionErr {
		t.Fatal("expected image not in collection error")
	}
}

func TestHandleInternalErrOnAddLabel(t *testing.T) {
	presenter := &FakePresenter{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{})
	repo := &FakeRepo{ErrOnAddLabel: true, Err: e.ErrInternal}
	itr := NewInteractor(repo, presenter, labelingService)
	itr.Execute(Request{Label: "a-label", ImageId: im.NewImageID(), Collection: "a-collection"})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error")
	}
}

func TestAssignLabelToImage(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	labelingService := lbi.NewLabelingService(&lbi.FakeRepo{})
	itr := NewInteractor(repo, presenter, labelingService)
	itr.Execute(Request{Label: "a-label", ImageId: im.NewImageID(), Collection: "a-collection"})
	if !presenter.GotSuccess || !repo.GotLabel || !(presenter.Got.Label == "a-label") {
		t.Fatal("expected to add one label")
	}
}
