package reader

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type FakeReadArtefactRepo struct {
	Data []byte
}

func (r *FakeReadArtefactRepo) Get(im.ImageId) ([]byte, error) {
	return r.Data, nil
}
