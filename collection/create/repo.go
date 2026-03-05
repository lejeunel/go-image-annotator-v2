package create

import (
	c "github.com/lejeunel/go-image-annotator-v2/collection"
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
		return c.ErrDuplicate
	}
	r.Got = req
	return nil
}
