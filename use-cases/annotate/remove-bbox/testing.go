package remove_bbox

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
)

type FakeRepo struct {
	Err error
	Got a.AnnotationId
}

func (r *FakeRepo) RemoveAnnotation(annotationId a.AnnotationId) error {
	if r.Err != nil {
		return r.Err
	}
	r.Got = annotationId
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
