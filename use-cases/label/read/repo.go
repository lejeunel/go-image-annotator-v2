package read

import (
	l "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Repo interface {
	Find(string) (*l.Label, error)
}

type FakeRepo struct {
	Label l.Label
	Err   error
}

func (r *FakeRepo) Find(name string) (*l.Label, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	if name == r.Label.Name {
		return &r.Label, nil
	}
	return nil, e.ErrNotFound

}
