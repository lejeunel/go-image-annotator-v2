package create

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"slices"
)

type CreateRepo interface {
	Create(CreateModel) error
}

type FakeCreateRepo struct {
	Names []string
	Got   CreateModel
}

func (r *FakeCreateRepo) Create(m CreateModel) error {
	if slices.Contains(r.Names, m.Name) {
		return e.ErrDuplicate
	}
	r.Got = m
	return nil
}

type FakeInternalErrCreateRepo struct{}

func (r *FakeInternalErrCreateRepo) Create(m CreateModel) error {
	return e.ErrInternal
}
