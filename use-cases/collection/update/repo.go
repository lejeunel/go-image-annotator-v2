package update

import (
	"slices"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type Repo interface {
	Update(Model) error
	Exists(string) (bool, error)
}

type FakeRepo struct {
	Names []string
	Got   Model
}

func (r *FakeRepo) Update(m Model) error {
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

func (r *FakeErrRepo) Update(m Model) error {
	return e.ErrInternal
}

func (r *FakeErrRepo) Exists(n string) (bool, error) {
	return false, r.err
}
