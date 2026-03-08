package read

import (
	c "github.com/lejeunel/go-image-annotator-v2/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type ListRepo interface {
	List(ListRequest) ([]*c.Collection, error)
}

type FakeListRepo struct {
	ReturnedSomething bool
}

func (r *FakeListRepo) List(req ListRequest) ([]*c.Collection, error) {
	r.ReturnedSomething = true
	return []*c.Collection{}, nil
}

type FakeInternalErrListRepo struct{}

func (r *FakeInternalErrListRepo) List(req ListRequest) ([]*c.Collection, error) {
	return nil, e.ErrInternal

}
