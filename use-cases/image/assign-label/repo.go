package assign_label

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type Repo interface {
	AddLabel(im.ImageId, clc.CollectionId, lbl.LabelId) error
}

type FakeRepo struct {
	Err           error
	ErrOnAddLabel bool
	GotLabel      bool
}

func (r *FakeRepo) AddLabel(im.ImageId, clc.CollectionId, lbl.LabelId) error {
	if r.ErrOnAddLabel {
		return r.Err
	}
	r.GotLabel = true
	return nil

}
