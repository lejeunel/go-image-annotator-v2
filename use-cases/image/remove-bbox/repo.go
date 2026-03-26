package remove_bbox

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
)

type Repo interface {
	RemoveAnnotation(a.AnnotationId) error
}

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
