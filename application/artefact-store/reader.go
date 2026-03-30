package artefact_store

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"io"
)
import ()

type ArtefactReadRepo interface {
	Get(im.ImageId) ([]byte, error)
}

type FromStoreImageReader struct {
	repo ArtefactReadRepo
	id   im.ImageId
}

func NewImageReader(id im.ImageId, repo ArtefactReadRepo) *FromStoreImageReader {
	return &FromStoreImageReader{repo: repo, id: id}
}
func (r FromStoreImageReader) Read(buf []byte) (int, error) {
	data, err := r.repo.Get(r.id)
	if err != nil {
		return 0, err
	}
	n := copy(buf, data)
	return n, io.EOF

}
