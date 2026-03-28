package read

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Repo interface {
	FindCollectionByName(string) (*clc.Collection, error)
}

type FakeRepo struct {
	Err        error
	Collection clc.Collection
}

func (r *FakeRepo) FindCollectionByName(name string) (*clc.Collection, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	if name == r.Collection.Name {
		return &r.Collection, nil
	}
	return nil, e.ErrNotFound

}
