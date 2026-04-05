package read_raw

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type FakeRepo struct {
	Err  error
	Data []byte
}

func (r *FakeRepo) Get(id im.ImageId) ([]byte, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	return r.Data, nil
}

type FakePresenter struct {
	Got            Response
	GotNotFoundErr bool
	GotSuccess     bool
	GotInternalErr bool
}

func (p *FakePresenter) ErrNotFound(err error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(err error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
