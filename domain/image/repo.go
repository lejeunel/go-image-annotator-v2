package image

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type CollectionRepo interface {
	FindCollectionByName(string) (*clc.Collection, error)
}

type AnnotationRepo interface {
	FindLabels(ImageId, clc.CollectionId) ([]*a.ImageLabel, error)
	FindBoundingBoxes(ImageId, clc.CollectionId) ([]*a.BoundingBox, error)
}

type ImageRepo interface {
	ImageExistsInCollection(ImageId, clc.CollectionId) (bool, error)
}

type FakeCollectionRepo struct {
	Err                 error
	MissingCollection   bool
	ErrOnFindCollection bool
	Collection          clc.Collection
}

type FakeAnnotationRepo struct {
	Err                    error
	ErrOnFindImageLabel    bool
	ErrOnFindBoundingBoxes bool
	Labels                 []*a.ImageLabel
	BoundingBoxes          []*a.BoundingBox
}

type FakeImageRepo struct {
	Err         error
	ErrOnExists bool
}

func (r *FakeImageRepo) ImageExistsInCollection(imageId ImageId, collectionId clc.CollectionId) (bool, error) {
	if r.ErrOnExists {
		return false, r.Err
	}
	return true, nil
}

func (r *FakeAnnotationRepo) FindBoundingBoxes(imageId ImageId, collectionId clc.CollectionId) ([]*a.BoundingBox, error) {
	if r.ErrOnFindBoundingBoxes {
		return nil, r.Err
	}
	if r.BoundingBoxes != nil {
		return r.BoundingBoxes, nil
	}
	return nil, nil
}

func (r *FakeAnnotationRepo) FindLabels(imageId ImageId, collectionId clc.CollectionId) ([]*a.ImageLabel, error) {
	if r.ErrOnFindImageLabel {
		return nil, r.Err
	}
	if r.Labels != nil {
		return r.Labels, nil
	}
	return nil, nil
}

func (r *FakeCollectionRepo) FindCollectionByName(name string) (*clc.Collection, error) {
	if r.MissingCollection {
		return nil, e.ErrNotFound
	}
	if r.ErrOnFindCollection {
		return nil, r.Err
	}
	return &r.Collection, nil
}
