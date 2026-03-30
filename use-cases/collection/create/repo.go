package create

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	"slices"
)

type CreateRepo interface {
	Create(clc.Collection) error
	Exists(string) (bool, error)
}

type FakeRepo struct {
	Err   error
	Names []string
	Got   clc.Collection
}

func (r *FakeRepo) Create(c clc.Collection) error {
	if r.Err != nil {
		return r.Err
	}

	r.Got = c
	return nil
}

func (r *FakeRepo) Exists(name string) (bool, error) {
	if slices.Contains(r.Names, name) {
		return true, nil
	}
	return false, nil
}
