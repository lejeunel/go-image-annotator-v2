package read_raw

import (
	"bytes"
	"io"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	t "github.com/lejeunel/go-image-annotator-v2/shared/testing"
)

type FakeRepo struct {
	Err  error
	Data []byte
}

func (r *FakeRepo) Get(id im.ImageId) (io.Reader, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	return bytes.NewBuffer(r.Data), nil
}

type FakePresenter struct {
	Got        Response
	GotSuccess bool
	t.TestingErrPresenter
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
