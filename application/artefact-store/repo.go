package artefact_store

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type ArtefactRepo interface {
	Store(im.ImageId, []byte) error
	Delete(im.ImageId) error
	Get(im.ImageId) ([]byte, error)
}

type FakeArtefactRepo struct {
	GotArtefact      bool
	Err              error
	NumDeletedImages int
	Data             []byte
}

func (r *FakeArtefactRepo) Store(im.ImageId, []byte) error {
	if r.Err != nil {
		return r.Err
	}
	r.GotArtefact = true
	return nil
}

func (r *FakeArtefactRepo) Delete(im.ImageId) error {
	r.NumDeletedImages += 1
	return nil
}

func (r *FakeArtefactRepo) Get(im.ImageId) ([]byte, error) {
	return r.Data, nil
}
