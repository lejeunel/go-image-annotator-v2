package image

import (
	"io"
)

type FakeImageStore struct {
	Err    error
	Got    BaseImage
	Return *Image
}

func (s *FakeImageStore) Find(baseImage BaseImage) (*Image, error) {
	if s.Err != nil {
		return nil, s.Err
	}
	if s.Return != nil {
		return s.Return, nil
	}
	return &Image{}, nil
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
