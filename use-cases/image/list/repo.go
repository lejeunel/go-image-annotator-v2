package list

import (
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type Repo interface {
	List(im.FilteringParams) ([]*im.BaseImage, error)
	Count(im.CountingParams) (*int64, error)
}

type FakeRepo struct {
	GotFilters im.FilteringParams
	Err        error
	Count_     int64
	ErrOnList  bool
	ErrOnCount bool
}

func (r *FakeRepo) List(f im.FilteringParams) ([]*im.BaseImage, error) {
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
func (r *FakeRepo) Count(f im.CountingParams) (*int64, error) {
	if r.ErrOnCount {
		return nil, r.Err
	}
	return &r.Count_, nil

}
