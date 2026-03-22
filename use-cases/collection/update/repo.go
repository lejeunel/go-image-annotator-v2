package update

import (
	"slices"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Repo interface {
	Update(UpdateModel) error
	Exists(string) (bool, error)
}

type FakeRepo struct {
	Names []string
	Got   UpdateModel
}

func (r *FakeRepo) Update(m UpdateModel) error {
	r.Got = m
	return nil
}
func (r *FakeRepo) Exists(n string) (bool, error) {
	if slices.Contains(r.Names, n) {
		return true, nil
	}
	return false, nil
}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Update(m UpdateModel) error {
	return e.ErrInternal
}

func (r *FakeErrRepo) Exists(n string) (bool, error) {
	return false, r.err
}
