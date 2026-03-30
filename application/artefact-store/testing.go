package artefact_store

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"io"
)

type FakeImageReader struct {
	Err  error
	Data []byte
}

func (r FakeImageReader) Read(buf []byte) (int, error) {
	if r.Err != nil {
		return 0, r.Err
	}
	n := copy(buf, r.Data)
	return n, io.EOF

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
