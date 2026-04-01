package list

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestHandleInternalErrOnCount(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnCount: true, Err: e.ErrInternal})
	itr.Execute(Request{}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatal("expected internal error, but got none")
	}
}

func TestHandleInternalErrOnList(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnList: true, Err: e.ErrInternal})
	itr.Execute(Request{}, p)
	if !p.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
	if p.GotSuccess {
		t.Fatalf("expected to get no success")
	}
}

func TestListCollection(t *testing.T) {
	count := int64(3)
	pageSize := 2
	page := int64(1)

	repo := &FakeRepo{Count_: count}
	p := &FakePresenter{}
	itr := NewInteractor(repo)
	req := Request{PageSize: pageSize, Page: page}
	itr.Execute(req, p)
	got := p.Got
	if len(got.Collections) != pageSize {
		t.Fatalf("expected to retrieve %v collections, got %v", pageSize, len(got.Collections))
	}
	if got.Pagination.TotalRecords != count {
		t.Fatalf("expected to retrieve count of %v, got %v", count, got.Pagination.TotalRecords)
	}
	if int(got.Pagination.TotalPages) != 2 {
		t.Fatalf("expected to retrieve total pages %v, got %v", 2, got.Pagination.TotalPages)
	}
	if got.Pagination.Page != page {
		t.Fatalf("expected to retrieve page %v, got %v", page, got.Pagination.Page)
	}
	if int(got.Pagination.PageSize) != pageSize {
		t.Fatalf("expected to retrieve page %v, got %v", pageSize, got.Pagination.Page)
	}
}
