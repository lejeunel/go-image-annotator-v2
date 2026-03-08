package create

import (
	"slices"
)

type CreateRepo interface {
	Create(CreateModel) error
	Exists(string) (bool, error)
}

type FakeCreateRepo struct {
	Names []string
	Got   CreateModel
}

func (r *FakeCreateRepo) Create(m CreateModel) error {
	r.Got = m
	return nil
}
func (r *FakeCreateRepo) Exists(name string) (bool, error) {
	if slices.Contains(r.Names, name) {
		return true, nil
	}
	return false, nil
}

type FakeErrCreateRepo struct {
	err error
}

func (r *FakeErrCreateRepo) Create(m CreateModel) error {
	return r.err
}

func (r *FakeErrCreateRepo) Exists(string) (bool, error) {
	return false, r.err
}
