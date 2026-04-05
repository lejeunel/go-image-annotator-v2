package modify_bbox

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
)

type FakeRepo struct {
	Err            error
	ErrOnUpdate    bool
	ErrOnFindLabel bool
	Got            a.BoundingBoxUpdatables
	Label          lbl.Label
}

func (r *FakeRepo) UpdateBoundingBox(id a.AnnotationId, u a.BoundingBoxUpdatables) error {
	if r.ErrOnUpdate {
		return r.Err
	}
	r.Got = u
	return nil
}

func (r *FakeRepo) FindLabelByName(name string) (*lbl.Label, error) {
	if r.ErrOnFindLabel {
		return nil, r.Err
	}
	return &r.Label, nil
}

type FakePresenter struct {
	GotNotFoundErr   bool
	GotInternalErr   bool
	GotValidationErr bool
	GotSuccess       bool
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

func (p *FakePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}
