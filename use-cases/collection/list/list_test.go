package list

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestHandleInternalErrOnCount(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnCount: true, Err: e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatal("expected internal error, but got none")
	}
}

func TestHandleInternalErrOnList(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnList: true, Err: e.ErrInternal}, presenter)
	itr.Execute(Request{})
	if !presenter.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
	if presenter.GotSuccess {
		t.Fatalf("expected to get no success")
	}
}

func TestListCollection(t *testing.T) {
	count := 3
	pageSize := 2
	page := 1

	repo := &FakeRepo{Count_: count}
	presenter := &FakePresenter{}
	itr := NewInteractor(repo, presenter)
	req := Request{PageSize: pageSize, Page: page}
	itr.Execute(req)
	got := presenter.Got
	if len(got.Collections) != pageSize {
		t.Fatalf("expected to retrieve %v collections, got %v", pageSize, len(got.Collections))
	}
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
