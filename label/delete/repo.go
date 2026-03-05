package delete

import (
	"slices"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type DeleteLabelRepo interface {
	Delete(r DeleteLabelModel) error
}

type FakeDeleteLabelRepo struct {
	ArePopulated []string
}

type FakeInternalErrDeleteLabelRepo struct{}

func (r *FakeInternalErrDeleteLabelRepo) Delete(m DeleteLabelModel) error {
	return e.ErrInternal
}

func (r *FakeDeleteLabelRepo) Delete(m DeleteLabelModel) error {
	if slices.Contains(r.ArePopulated, m.Name) {
		return e.ErrDependency
	}

	return nil
}
