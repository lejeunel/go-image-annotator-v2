package create

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"slices"
)

type CreateCollectionRepo interface {
	Create(CreateCollectionRequest) error
}

type FakeCreateCollectionRepo struct {
	Names []string
	Got   CreateCollectionRequest
}

func (r *FakeCreateCollectionRepo) Create(req CreateCollectionRequest) error {
	if slices.Contains(r.Names, req.Name) {
		return e.ErrDuplicate
	}
	r.Got = req
	return nil
}

type FakeInternalErrCreateCollectionRepo struct {
}

func (r *FakeInternalErrCreateCollectionRepo) Create(req CreateCollectionRequest) error {
	return e.ErrInternal
}
