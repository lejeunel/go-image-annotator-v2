package read

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	l "github.com/lejeunel/go-image-annotator-v2/label"
)

type ListRepo interface {
	List(ListRequest) ([]*l.Label, error)
}

type FakeListRepo struct {
	ReturnedSomething bool
}

func (r *FakeListRepo) List(req ListRequest) ([]*l.Label, error) {
	r.ReturnedSomething = true
	return []*l.Label{}, nil

}

type FakeInternalErrListRepo struct{}

func (r *FakeInternalErrListRepo) List(req ListRequest) ([]*l.Label, error) {
	return nil, e.ErrInternal

}
