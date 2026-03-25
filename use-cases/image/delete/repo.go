package delete

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type Repo interface {
	Delete(im.ImageId) error
	RemoveAnnotation(im.ImageId, clc.CollectionId, a.AnnotationId) error
}

type FakeRepo struct {
	Err                   error
	DeletedId             im.ImageId
	ErrOnRemoveAnnotation bool
}

func (r *FakeRepo) Delete(id im.ImageId) error {
	if r.Err != nil {
		return r.Err
	}
	r.DeletedId = id
	return nil
}

func (r *FakeRepo) RemoveAnnotation(imageId im.ImageId, collectionId clc.CollectionId, annotationId a.AnnotationId) error {
	if r.ErrOnRemoveAnnotation {
		return r.Err
	}
	return nil
}
