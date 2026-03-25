package read

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Repo interface {
	Find(string) (*clc.Collection, error)
}

type FakeReadRepo struct {
	Collection clc.Collection
}

func (r *FakeReadRepo) Find(name string) (*clc.Collection, error) {

	if name == r.Collection.Name {
		return &r.Collection, nil
	}
	return nil, e.ErrNotFound

}

type FakeInternalErrReadRepo struct{}

func (r *FakeInternalErrReadRepo) Find(name string) (*clc.Collection, error) {
	return nil, e.ErrInternal

}
