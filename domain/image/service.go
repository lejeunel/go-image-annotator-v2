package image

type ImageService interface {
	Find(BaseImage) (*Image, error)
}

type FakeImageService struct {
	Err error
}

func (s *FakeImageService) Find(baseImage BaseImage) (*Image, error) {
	if s.Err != nil {
		return nil, s.Err
	}
	return &Image{}, nil
}
