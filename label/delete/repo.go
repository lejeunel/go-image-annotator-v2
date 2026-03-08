package delete

import (
	"slices"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type DeleteRepo interface {
	Delete(r DeleteModel) error
}

type FakeDeleteRepo struct {
	ArePopulated []string
}

type FakeInternalErrDeleteRepo struct{}

func (r *FakeInternalErrDeleteRepo) Delete(m DeleteModel) error {
	return e.ErrInternal
}

func (r *FakeDeleteRepo) Delete(m DeleteModel) error {
	if slices.Contains(r.ArePopulated, m.Name) {
		return e.ErrDependency
	}

	return nil
}
