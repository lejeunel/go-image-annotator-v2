package create

import (
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	"slices"
)

type Repo interface {
	Create(lbl.Label) error
	LabelWithNameExists(string) (bool, error)
}

type FakeRepo struct {
	Err   error
	Names []string
	Got   lbl.Label
}

func (r *FakeRepo) Create(l lbl.Label) error {
	if r.Err != nil {
		return r.Err
	}
	r.Got = l
	return nil
}
func (r *FakeRepo) LabelWithNameExists(name string) (bool, error) {
	if slices.Contains(r.Names, name) {
		return true, nil
	}
	return false, nil
}
