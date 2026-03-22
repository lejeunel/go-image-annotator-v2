package create

import (
	"slices"
)

type CreateRepo interface {
	Create(Model) error
	Exists(string) (bool, error)
}

type FakeRepo struct {
	Names []string
	Got   Model
}

func (r *FakeRepo) Create(m Model) error {
	r.Got = m
	return nil
}

func (r *FakeRepo) Exists(name string) (bool, error) {
	if slices.Contains(r.Names, name) {
		return true, nil
	}
	return false, nil
}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Create(m Model) error {
	return r.err
}

func (r *FakeErrRepo) Exists(string) (bool, error) {
	return false, r.err
}
