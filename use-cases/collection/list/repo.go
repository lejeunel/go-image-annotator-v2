package list

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
)

type Repo interface {
	List(Request) ([]*clc.Collection, error)
}

type FakeRepo struct {
	ReturnedSomething bool
}

func (r *FakeRepo) List(req Request) ([]*clc.Collection, error) {
	r.ReturnedSomething = true
	return []*clc.Collection{}, nil
}

type FakeErrListRepo struct {
	err error
}

func (r *FakeErrListRepo) List(req Request) ([]*clc.Collection, error) {
	return nil, r.err

}
