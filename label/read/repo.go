package read

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	l "github.com/lejeunel/go-image-annotator-v2/label"
)

type ReadRepo interface {
	Find(string) (*l.Label, error)
}

type FakeReadRepo struct {
	Label l.Label
}

func (r *FakeReadRepo) Find(name string) (*l.Label, error) {

	if name == r.Label.Name {
		return &r.Label, nil
	}
	return nil, e.ErrNotFound

}

type FakeInternalErrReadRepo struct{}

func (r *FakeInternalErrReadRepo) Find(name string) (*l.Label, error) {
	return nil, e.ErrInternal

}
