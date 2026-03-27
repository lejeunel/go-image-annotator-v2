package list

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestHandleInternalErrOnList(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnList: true, Err: e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error, but got none")
	}
}

func TestHandleInternalErrOnCount(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnCount: true, Err: e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error, but got none")
	}
}

func TestListLabel(t *testing.T) {
	count := 3
	pageSize := 2
	page := 1
	repo := &FakeRepo{Count_: count}
	presenter := &FakePresenter{}
	itr := NewInteractor(repo, presenter)
	itr.Execute(Request{PageSize: pageSize, Page: page})
	if len(presenter.Got.Labels) != pageSize {
		t.Fatalf("expected to retrieve %v labels, got %v", pageSize, len(presenter.Got.Labels))
	}
	got := presenter.Got
	if int(got.Pagination.TotalRecords) != count {
		t.Fatalf("expected to retrieve count of %v, got %v", count, got.Pagination.TotalRecords)
	}
	if int(got.Pagination.TotalPages) != 2 {
		t.Fatalf("expected to retrieve total pages %v, got %v", 2, got.Pagination.TotalPages)
	}
	if int(got.Pagination.Page) != page {
		t.Fatalf("expected to retrieve page %v, got %v", page, got.Pagination.Page)
	}
	if int(got.Pagination.PageSize) != pageSize {
		t.Fatalf("expected to retrieve page %v, got %v", pageSize, got.Pagination.Page)
	}
}
