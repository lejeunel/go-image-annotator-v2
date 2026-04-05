package read

import (
	l "github.com/lejeunel/go-image-annotator-v2/entities/label"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type FakeRepo struct {
	Label l.Label
	Err   error
}

func (r *FakeRepo) FindLabelByName(name string) (*l.Label, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	if name == r.Label.Name {
		return &r.Label, nil
	}
	return nil, e.ErrNotFound

}

type FakePresenter struct {
	Got            Response
	GotNotFoundErr bool
	GotInternalErr bool
	GotSuccess     bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
