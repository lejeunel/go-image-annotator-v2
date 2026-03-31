package unassign_label

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
)

type Repo interface {
	RemoveLabel(im.ImageId, clc.CollectionId, lbl.LabelId) error
}

type FakeRepo struct {
	Err              error
	ErrOnRemoveLabel bool
	RemovedLabel     bool
}

func (r *FakeRepo) RemoveLabel(im.ImageId, clc.CollectionId, lbl.LabelId) error {
	if r.ErrOnRemoveLabel {
		return r.Err
	}
	r.RemovedLabel = true
	return nil

}
