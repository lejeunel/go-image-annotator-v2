package delete

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type Repo interface {
	RemoveImageFromCollection(im.ImageId, clc.CollectionId) error
	RemoveAnnotation(im.ImageId, clc.CollectionId, a.AnnotationId) error
}

type FakeRepo struct {
	Err                   error
	RemovedImageId        im.ImageId
	ErrOnRemoveImage      bool
	ErrOnRemoveAnnotation bool
}

func (r *FakeRepo) RemoveImageFromCollection(imageId im.ImageId, collectionId clc.CollectionId) error {
	if r.Err != nil {
		return r.Err
	}
	r.RemovedImageId = imageId
	return nil
}

func (r *FakeRepo) RemoveAnnotation(imageId im.ImageId, collectionId clc.CollectionId, annotationId a.AnnotationId) error {
	if r.ErrOnRemoveAnnotation {
		return r.Err
	}
	return nil
}
