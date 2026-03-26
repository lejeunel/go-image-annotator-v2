package modify_bbox

import (
	"testing"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{Err: e.ErrNotFound, ErrOnFindLabel: true})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestInternalErrOnFindLabelShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{Err: e.ErrInternal, ErrOnFindLabel: true})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected internal error")
	}
}

func TestValidationErrShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{})
	itr.Execute(Request{Xc: 1, Yc: 1, Width: -999, Height: 1})
	if !presenter.GotValidationErr || presenter.GotSuccess {
		t.Fatalf("expected validation error")
	}
}

func TestNotFoundErrOnUpdateShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{ErrOnUpdate: true, Err: e.ErrNotFound})
	itr.Execute(Request{Xc: 1, Yc: 1, Width: 1, Height: 1})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected not found error")
	}
}

func TestInternalErrOnUpdateShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(presenter, &FakeRepo{ErrOnUpdate: true, Err: e.ErrInternal})
	itr.Execute(Request{Xc: 1, Yc: 1, Width: 1, Height: 1})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected internal error")
	}
}

func TestUpdate(t *testing.T) {
	presenter := &FakePresenter{}
	label := lbl.NewLabel(lbl.NewLabelID(), "a-label")
	repo := &FakeRepo{Label: *label}
	itr := NewInteractor(presenter, repo)
	annotationId := a.NewAnnotationId()
	r := Request{AnnotationId: annotationId, Xc: 1, Yc: 1, Width: 1, Height: 1}
	itr.Execute(r)
	got := repo.Got
	want := Updatables{LabelId: label.Id,
		AnnotationId: annotationId, Xc: r.Xc,
		Yc: r.Yc, Width: r.Width, Height: r.Height}
	if !presenter.GotSuccess {
		t.Fatalf("expected success")
	}
	if got != want {
		t.Fatalf("expected to update with %+v, got %+v", want, got)
	}
}
