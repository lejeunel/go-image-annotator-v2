package list

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
)

type Repo interface {
	List(Request) ([]*clc.Collection, error)
	Count() (int64, error)
}

type FakeRepo struct {
	Err        error
	ErrOnCount bool
	ErrOnList  bool
	Count_     int
}

func (r *FakeRepo) Count() (int64, error) {
	if r.ErrOnCount {
		return 0, r.Err
	}
	return int64(r.Count_), nil
}

func (r *FakeRepo) List(req Request) ([]*clc.Collection, error) {
	if r.ErrOnList {
		return nil, r.Err
	}

	result := []*clc.Collection{}
	for range req.PageSize {
		result = append(result, clc.NewCollection(clc.NewCollectionId(), "a-collection"))
	}
	return result, nil
}
