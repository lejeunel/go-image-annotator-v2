package image_store

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"io"
)

type FakeImageStore struct {
	Err    error
	Got    im.BaseImage
	Return *im.Image
}

func (s *FakeImageStore) Find(baseImage im.BaseImage) (*im.Image, error) {
	if s.Err != nil {
		return nil, s.Err
	}
	if s.Return != nil {
		return s.Return, nil
	}
	return &im.Image{}, nil
}

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
