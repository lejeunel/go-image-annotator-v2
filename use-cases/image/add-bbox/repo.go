package add_bbox

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type Repo interface {
	AddBoundingBox(im.ImageId, clc.CollectionId, a.BoundingBox) error
	FindLabelByName(string) (*lbl.Label, error)
}

type FakeRepo struct {
	Err             error
	ErrOnAdd        bool
	ErrOnFindLabel  bool
	GotImageId      im.ImageId
	GotCollectionId clc.CollectionId
	GotBox          a.BoundingBox
}

func (r *FakeRepo) AddBoundingBox(imageId im.ImageId, collectionId clc.CollectionId, box a.BoundingBox) error {
	if r.ErrOnAdd {
		return r.Err
	}
	r.GotImageId = imageId
	r.GotCollectionId = collectionId
	r.GotBox = box
	return nil
}

func (r *FakeRepo) FindLabelByName(name string) (*lbl.Label, error) {
	if r.ErrOnFindLabel {
		return nil, r.Err
	}
	return lbl.NewLabel(lbl.NewLabelID(), name), nil
}
