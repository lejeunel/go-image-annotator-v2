package read

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	c "github.com/lejeunel/go-image-annotator-v2/label"
)

type ReadRepo interface {
	Find(string) (*c.Label, error)
}

type FakeReadRepo struct {
	Label c.Label
}

func (r *FakeReadRepo) Find(name string) (*c.Label, error) {

	if name == r.Label.Name {
		return &r.Label, nil
	}
	return nil, e.ErrNotFound

}

type FakeInternalErrReadRepo struct{}

func (r *FakeInternalErrReadRepo) Find(name string) (*c.Label, error) {
	return nil, e.ErrInternal

}
