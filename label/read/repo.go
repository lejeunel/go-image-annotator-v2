package read

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	l "github.com/lejeunel/go-image-annotator-v2/label"
)

type Repo interface {
	Find(string) (*l.Label, error)
}

type FakeRepo struct {
	Label l.Label
}

func (r *FakeRepo) Find(name string) (*l.Label, error) {

	if name == r.Label.Name {
		return &r.Label, nil
	}
	return nil, e.ErrNotFound

}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Find(name string) (*l.Label, error) {
	return nil, r.err

}
