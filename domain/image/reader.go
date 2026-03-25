package image

import (
	"io"
)

type ArtefactReadRepo interface {
	Get(ImageId) ([]byte, error)
}

type ImageReader interface {
	Read([]byte) (int, error)
}

type FromStoreImageReader struct {
	repo ArtefactReadRepo
	id   ImageId
}

func NewImageReader(id ImageId, repo ArtefactReadRepo) *FromStoreImageReader {
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
