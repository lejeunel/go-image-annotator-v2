package update

import (
	c "github.com/lejeunel/go-image-annotator-v2/collection"
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
	if slices.Contains(r.Names, req.Name) {
		return c.ErrDuplicate
	}
	r.Got = req
	return nil
}
