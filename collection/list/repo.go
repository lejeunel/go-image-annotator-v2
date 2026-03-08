package list

import (
	c "github.com/lejeunel/go-image-annotator-v2/collection"
)

type Repo interface {
	List(Request) ([]*c.Collection, error)
}

type FakeRepo struct {
	ReturnedSomething bool
}

func (r *FakeRepo) List(req Request) ([]*c.Collection, error) {
	r.ReturnedSomething = true
	return []*c.Collection{}, nil
}

type FakeErrListRepo struct {
	err error
}

func (r *FakeErrListRepo) List(req Request) ([]*c.Collection, error) {
	return nil, r.err

}
