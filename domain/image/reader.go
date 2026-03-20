package image

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/artefact"
)

type ImageReader interface {
	Read() (*RawImage, error)
}

type FakeImageReader struct {
	Err error
}

func (r FakeImageReader) Read() (*RawImage, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	return &RawImage{ArtefactId: a.NewArtefactId()}, nil

}
