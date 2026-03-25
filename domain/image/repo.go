package image

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Repo interface {
	FindLabels(ImageId, clc.CollectionId) ([]*a.ImageLabel, error)
	FindBoundingBoxes(ImageId, clc.CollectionId) ([]*a.BoundingBox, error)
	FindCollectionByName(string) (*clc.Collection, error)
	ImageExistsInCollection(ImageId, string) (bool, error)
}

type FakeRepo struct {
	Err                    error
	ErrOnFindLabel         bool
	ErrOnFindBoundingBoxes bool
	ErrOnExists            bool
	MissingCollection      bool
	Collection             *clc.Collection
	Labels                 []*a.ImageLabel
	BoundingBoxes          []*a.BoundingBox
}

func (r *FakeRepo) ImageExistsInCollection(imageId ImageId, collectionName string) (bool, error) {
	if r.ErrOnExists {
		return false, r.Err
	}
	return true, nil
}

func (r *FakeRepo) FindBoundingBoxes(imageId ImageId, collectionId clc.CollectionId) ([]*a.BoundingBox, error) {
	if r.ErrOnFindBoundingBoxes {
		return nil, r.Err
	}
	if r.BoundingBoxes != nil {
		return r.BoundingBoxes, nil
	}
	return nil, nil
}

func (r *FakeRepo) FindLabels(imageId ImageId, collectionId clc.CollectionId) ([]*a.ImageLabel, error) {
	if r.ErrOnFindLabel {
		return nil, r.Err
	}
	if r.Labels != nil {
		return r.Labels, nil
	}
	return nil, nil
}

func (r *FakeRepo) FindCollectionByName(name string) (*clc.Collection, error) {
	if r.MissingCollection {
		return nil, e.ErrNotFound
	}
	if r.Collection != nil {
		return r.Collection, nil
	}
	return nil, nil
}
