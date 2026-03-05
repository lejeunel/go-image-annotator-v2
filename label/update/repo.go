package update

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"slices"
)

type UpdateRepo interface {
	Update(UpdateModel) error
}

type FakeUpdateRepo struct {
	Names []string
	Got   UpdateModel
}

func (r *FakeUpdateRepo) Update(m UpdateModel) error {
	if !slices.Contains(r.Names, m.Name) {
		return e.ErrNotFound
	}
	if slices.Contains(r.Names, m.NewName) {
		return e.ErrDuplicate
	}
	r.Got = m
	return nil
}

type FakeInternalErrUpdateRepo struct{}

func (r *FakeInternalErrUpdateRepo) Update(m UpdateModel) error {
	return e.ErrInternal
}
