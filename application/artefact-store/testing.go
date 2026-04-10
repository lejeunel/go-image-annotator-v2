package artefact_store

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type FakeArtefactRepo struct {
	GotArtefact      bool
	Err              error
	NumDeletedImages int
	Data             []byte
	GotData          []byte
}

func (r *FakeArtefactRepo) Store(id im.ImageId, data []byte) error {
	if r.Err != nil {
		return r.Err
	}
	r.GotArtefact = true
	r.GotData = data
	return nil
}

func (r *FakeArtefactRepo) Delete(im.ImageId) error {
	r.NumDeletedImages += 1
	return nil
}

func (r *FakeArtefactRepo) Get(im.ImageId) ([]byte, error) {
	return r.Data, nil
}
