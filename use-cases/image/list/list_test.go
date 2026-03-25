package list

import (
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestHandleNotFoundErrOnList(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnList: true, Err: e.ErrNotFound},
		presenter, &im.FakeImageStore{})
	itr.Execute(Request{})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestHandleInternalErrOnList(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnList: true, Err: e.ErrInternal},
		presenter, &im.FakeImageStore{})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestHandleInternalErrOnImageBuild(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, presenter, &im.FakeImageStore{Err: e.ErrInternal})
	itr.Execute(Request{PageSize: 1})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestHandleInternalErrOnCount(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnCount: true, Err: e.ErrInternal}, presenter, &im.FakeImageStore{})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestListImages(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, presenter, &im.FakeImageStore{})
	r := Request{Page: 1, PageSize: 2}
	itr.Execute(r)
	if !presenter.GotSuccess || (len(presenter.Got.Images) != r.PageSize) {
		t.Fatalf("expected to list images")
	}
}

func TestPaginationMetaData(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{Count_: 10}
	itr := NewInteractor(repo, presenter, &im.FakeImageStore{})
	r := Request{Page: 1, PageSize: 2}
	itr.Execute(r)
	pg := presenter.Got.Pagination
	if !(pg.Page == r.Page) || !(pg.PageSize == r.PageSize) || !(pg.Total == 10) || !(pg.TotalPages == 5) {
		t.Fatalf("expected pagination meta-data with page 1, page size 2, total 1, and total pages 5, got %+v", pg)
	}
}

func TestQueryCorrectPagination(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{Count_: 10}
	itr := NewInteractor(repo, presenter, &im.FakeImageStore{})
	r := Request{Page: 1, PageSize: 2}
	itr.Execute(r)
	f := repo.GotFilters
	if !(f.Page == r.Page) || !(f.PageSize == r.PageSize) {
		t.Fatalf("expected to query repo with page %v and page size %v, got %v and %v",
			r.Page, r.PageSize, repo.GotFilters.Page, repo.GotFilters.PageSize)

	}
}
