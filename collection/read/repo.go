package read

import (
	c "github.com/lejeunel/go-image-annotator-v2/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type ReadRepo interface {
	Find(string) (*c.Collection, error)
}

type FakeReadRepo struct {
	Collection c.Collection
}

func (r *FakeReadRepo) Find(name string) (*c.Collection, error) {

	if name == r.Collection.Name {
		return &r.Collection, nil
	}
	return nil, e.ErrNotFound

}

type FakeInternalErrReadRepo struct{}

func (r *FakeInternalErrReadRepo) Find(name string) (*c.Collection, error) {
	return nil, e.ErrInternal

}
