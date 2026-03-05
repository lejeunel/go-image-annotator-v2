package update

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"slices"
)

type UpdateCollectionRepo interface {
	Update(UpdateCollectionRequest) error
}

type FakeUpdateCollectionRepo struct {
	Names []string
	Got   UpdateCollectionRequest
}

func (r *FakeUpdateCollectionRepo) Update(req UpdateCollectionRequest) error {
	if !slices.Contains(r.Names, req.Name) {
		return e.ErrNotFound
	}
	if slices.Contains(r.Names, req.NewName) {
		return e.ErrDuplicate
	}
	r.Got = req
	return nil
}

type FakeInternalErrUpdateCollectionRepo struct{}

func (r *FakeInternalErrUpdateCollectionRepo) Update(req UpdateCollectionRequest) error {
	return e.ErrInternal
}
