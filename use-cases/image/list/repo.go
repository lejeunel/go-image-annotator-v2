package list

import (
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type FilteringParams struct {
	Collection *string
	PageSize   int
	Page       int
}

type Repo interface {
	List(FilteringParams) ([]*im.BaseImage, error)
	Count(FilteringParams) (*int, error)
}

type FakeRepo struct {
	GotFilters FilteringParams
	Err        error
	Count_     int
	ErrOnList  bool
	ErrOnCount bool
}

func (r *FakeRepo) List(f FilteringParams) ([]*im.BaseImage, error) {
	if r.ErrOnList {
		return nil, r.Err
	}

	r.GotFilters = f

	result := []*im.BaseImage{}
	collectionName := "a-collection"
	for range f.PageSize {
		result = append(result, &im.BaseImage{Collection: collectionName, ImageId: im.NewImageId()})
	}

	return result, nil

}
func (r *FakeRepo) Count(f FilteringParams) (*int, error) {
	if r.ErrOnCount {
		return nil, r.Err
	}
	return &r.Count_, nil

}
