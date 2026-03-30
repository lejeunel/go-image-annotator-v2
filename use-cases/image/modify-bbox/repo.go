package modify_bbox

import (
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
)

type Repo interface {
	FindLabelByName(string) (*lbl.Label, error)
	UpdateBoundingBox(Updatables) error
}

type FakeRepo struct {
	Err            error
	ErrOnUpdate    bool
	ErrOnFindLabel bool
	Got            Updatables
	Label          lbl.Label
}

func (r *FakeRepo) UpdateBoundingBox(updatables Updatables) error {
	if r.ErrOnUpdate {
		return r.Err
	}
	r.Got = updatables
	return nil
}

func (r *FakeRepo) FindLabelByName(name string) (*lbl.Label, error) {
	if r.ErrOnFindLabel {
		return nil, r.Err
	}
	return &r.Label, nil
}
