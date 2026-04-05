package unassign_label

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
)

type FakeRepo struct {
	Err              error
	ErrOnRemoveLabel bool
	RemovedLabel     bool
}

func (r *FakeRepo) RemoveImageLabel(im.ImageId, clc.CollectionId, lbl.LabelId) error {
	if r.ErrOnRemoveLabel {
		return r.Err
	}
	r.RemovedLabel = true
	return nil

}

type FakePresenter struct {
	GotSuccess       bool
	GotNotFoundErr   bool
	GotInternalErr   bool
	GotDependencyErr bool
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrDependency(error) {
	p.GotDependencyErr = true
}

func (p *FakePresenter) Success(Response) {
	p.GotSuccess = true
}
