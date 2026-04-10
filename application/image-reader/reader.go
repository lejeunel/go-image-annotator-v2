package reader

import (
	ast "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"io"
)

type ImageReader struct {
	repo ast.ArtefactReadRepo
	id   im.ImageId

	data []byte
	pos  int
}

func NewImageReader(id im.ImageId, repo ast.ArtefactReadRepo) *ImageReader {
	return &ImageReader{repo: repo, id: id}
}
func (r *ImageReader) Read(buf []byte) (int, error) {
	if r.data == nil {
		data, err := r.repo.Get(r.id)
		if err != nil {
			return 0, err
		}
		r.data = data
	}

	if r.pos >= len(r.data) {
		return 0, io.EOF
	}

	n := copy(buf, r.data[r.pos:])
	r.pos += n

	return n, nil
}
