package ingest

import (
	an "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type FakeRepo struct {
	GotImage              bool
	Err                   error
	ErrOnFindCollection   bool
	ErrOnLabelExists      bool
	ErrOnIngest           bool
	ErrOnAddLabel         bool
	ErrOnFindHash         bool
	ErrOnAddBoundingBox   bool
	ErrOnDeleteImage      bool
	MissingCollection     bool
	MissingLabel          bool
	HashAlreadyExists     bool
	NumLabelsAdded        int
	NumBoundingboxesAdded int
	NumDeletedImages      int
}

func (r *FakeRepo) FindCollectionByName(name string) (*clc.Collection, error) {
	if r.ErrOnFindCollection {
		return nil, r.Err
	}
	if r.MissingCollection {
		return nil, e.ErrNotFound
	}
	c := clc.NewCollection(clc.NewCollectionId(), "a-collection")
	return c, nil
}

func (r *FakeRepo) FindLabelByName(name string) (*lbl.Label, error) {
	if r.ErrOnLabelExists {
		return nil, r.Err
	}
	if r.MissingLabel {
		return nil, e.ErrNotFound
	}
	return lbl.NewLabel(lbl.NewLabelId(), name), nil
}

func (r *FakeRepo) FindImageByHash(hash string) (*im.Image, error) {
	if r.ErrOnFindHash {
		return nil, r.Err
	}
	if r.HashAlreadyExists {
		return im.NewImage(im.NewImageId(), *clc.NewCollection(clc.NewCollectionId(), "a-collection")), nil
	}
	return nil, e.ErrNotFound
}

func (r *FakeRepo) IngestImage(im.ImageId, clc.CollectionId) error {
	if r.ErrOnIngest {
		return r.Err
	}
	return nil
}

func (r *FakeRepo) AddLabelToImage(im.ImageId, clc.CollectionId, lbl.LabelId) error {
	if r.ErrOnAddLabel {
		return r.Err
	}
	r.NumLabelsAdded += 1
	return nil
}

func (r *FakeRepo) AddBoundingBoxToImage(im.ImageId, clc.CollectionId, an.BoundingBox) error {
	if r.ErrOnAddBoundingBox {
		return r.Err
	}
	r.NumBoundingboxesAdded += 1
	return nil
}

func (r *FakeRepo) DeleteImage(im.ImageId) error {
	if r.ErrOnDeleteImage {
		return r.Err
	}
	r.NumDeletedImages += 1
	return nil
}
