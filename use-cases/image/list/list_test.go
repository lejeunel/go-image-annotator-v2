package list

import (
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func TestNonExistingCollectionShouldFail(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{NonExistingCollection: true},
		presenter, &im.FakeImageService{})
	collectionName := "a-collection"
	itr.Execute(Request{CollectionName: &collectionName})
	if !presenter.GotNotFoundErr || presenter.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestHandleInternalErrOnFindCollection(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnFindCollection: true, Err: e.ErrInternal},
		presenter, &im.FakeImageService{})
	collectionName := "a-collection"
	itr.Execute(Request{CollectionName: &collectionName})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestHandleInternalErrOnList(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnList: true, Err: e.ErrInternal},
		presenter, &im.FakeImageService{})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestHandleInternalErrOnImageBuild(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{}, presenter, &im.FakeImageService{Err: e.ErrInternal})
	itr.Execute(Request{PageSize: 1})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestHandleInternalErrOnCount(t *testing.T) {
	presenter := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{ErrOnCount: true, Err: e.ErrInternal}, presenter, &im.FakeImageService{})
	itr.Execute(Request{})
	if !presenter.GotInternalErr || presenter.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestListImages(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{}
	itr := NewInteractor(repo, presenter, &im.FakeImageService{})
	r := Request{Page: 1, PageSize: 2}
	itr.Execute(r)
	if !presenter.GotSuccess || (len(presenter.Got.Images) != r.PageSize) {
		t.Fatalf("expected to list images")
	}
}

func TestPaginationMetaData(t *testing.T) {
	presenter := &FakePresenter{}
	repo := &FakeRepo{Count_: 10}
	itr := NewInteractor(repo, presenter, &im.FakeImageService{})
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
	itr := NewInteractor(repo, presenter, &im.FakeImageService{})
	r := Request{Page: 1, PageSize: 2}
	itr.Execute(r)
	f := repo.GotFilters
	if !(f.Page == r.Page) || !(f.PageSize == r.PageSize) {
		t.Fatalf("expected to query repo with page %v and page size %v, got %v and %v",
			r.Page, r.PageSize, repo.GotFilters.Page, repo.GotFilters.PageSize)

	}
}
