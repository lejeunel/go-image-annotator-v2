package read_raw

import (
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
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
