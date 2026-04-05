package delete

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

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

type FakePresenter struct {
	GotNotFoundErr bool
	GotInternalErr bool
	GotSuccess     bool
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) Success(Response) {
	p.GotSuccess = true
}
