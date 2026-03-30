package assign_label

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Repo interface {
	AddLabel(im.ImageId, clc.CollectionId, lbl.LabelId) error
	FindLabel(string) (*lbl.Label, error)
}

type FakeRepo struct {
	Err            error
	ErrOnAddLabel  bool
	ErrOnFindLabel bool
	MissingLabel   bool

	AddedLabelId        lbl.LabelId
	AddedOnImageId      im.ImageId
	AddedOnCollectionId clc.CollectionId

	ReturnLabel lbl.Label
}

func (r *FakeRepo) FindLabel(string) (*lbl.Label, error) {
	if r.MissingLabel {
		return nil, e.ErrNotFound
	}
	if r.ErrOnFindLabel {
		return nil, r.Err
	}
	return &r.ReturnLabel, nil
}

func (r *FakeRepo) AddLabel(imageId im.ImageId, collectionId clc.CollectionId, labelId lbl.LabelId) error {
	if r.ErrOnAddLabel {
		return r.Err
	}
	r.AddedLabelId = labelId
	r.AddedOnImageId = imageId
	r.AddedOnCollectionId = collectionId
	return nil

}
