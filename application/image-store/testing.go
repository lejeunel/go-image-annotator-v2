package image_store

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

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
	Err           error
	ErrOnExists   bool
	ErrOnMIMEType bool
	MIMEType_     string
}

func (r *FakeImageRepo) ImageExistsInCollection(imageId im.ImageId, collectionId clc.CollectionId) (bool, error) {
	if r.ErrOnExists {
		return false, r.Err
	}
	return true, nil
}
func (r *FakeImageRepo) MIMEType(imageId im.ImageId) (*string, error) {
	if r.ErrOnMIMEType {
		return nil, r.Err
	}
	return &r.MIMEType_, nil
}

func (r *FakeAnnotationRepo) FindBoundingBoxes(imageId im.ImageId, collectionId clc.CollectionId) ([]*a.BoundingBox, error) {
	if r.ErrOnFindBoundingBoxes {
		return nil, r.Err
	}
	if r.BoundingBoxes != nil {
		return r.BoundingBoxes, nil
	}
	return nil, nil
}

func (r *FakeAnnotationRepo) FindImageLabels(imageId im.ImageId, collectionId clc.CollectionId) ([]*a.ImageLabel, error) {
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

type FakeImageStore struct {
	Err    error
	Return *im.Image
}

func (s *FakeImageStore) Find(baseImage im.BaseImage) (*im.Image, error) {
	if s.Err != nil {
		return nil, s.Err
	}
	if s.Return != nil {
		return s.Return, nil
	}
	return &im.Image{}, nil
}
