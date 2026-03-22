package list

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type FilteringParams struct {
	CollectionId *clc.CollectionId
	PageSize     int
	Page         int
}

type Repo interface {
	List(FilteringParams) ([]*im.BaseImage, error)
	Count(FilteringParams) (*int, error)
	FindCollectionIdByName(string) (*clc.CollectionId, error)
}

type FakeRepo struct {
	GotFilters            FilteringParams
	Err                   error
	Count_                int
	NonExistingCollection bool
	ErrOnFindCollection   bool
	ErrOnList             bool
	ErrOnCount            bool
}

func (r *FakeRepo) List(f FilteringParams) ([]*im.BaseImage, error) {
	if r.ErrOnList {
		return nil, r.Err
	}

	r.GotFilters = f

	result := []*im.BaseImage{}
	for range f.PageSize {
		result = append(result, &im.BaseImage{CollectionId: clc.NewCollectionID(), ImageId: im.NewImageID()})
	}

	return result, nil

}
func (r *FakeRepo) FindCollectionIdByName(string) (*clc.CollectionId, error) {
	if r.NonExistingCollection {
		return nil, e.ErrNotFound
	}
	if r.ErrOnFindCollection {
		return nil, r.Err
	}
	return &clc.CollectionId{}, nil
}
func (r *FakeRepo) Count(f FilteringParams) (*int, error) {
	if r.ErrOnCount {
		return nil, r.Err
	}
	return &r.Count_, nil

}
