package list

import (
	l "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type Repo interface {
	List(Request) ([]*l.Label, error)
}

type FakeRepo struct {
	ReturnedSomething bool
}

func (r *FakeRepo) List(req Request) ([]*l.Label, error) {
	r.ReturnedSomething = true
	return []*l.Label{}, nil

}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) List(req Request) ([]*l.Label, error) {
	return nil, r.err

}
