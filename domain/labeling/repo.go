package labeling

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type Repo interface {
	FindLabelByName(string) (*lbl.Label, error)
	FindCollectionByName(string) (*clc.Collection, error)
	ImageIsInCollection(im.ImageId, clc.CollectionId) (bool, error)
}

type FakeRepo struct {
	Err                      error
	ErrOnFindLabel           bool
	ErrOnFindCollection      bool
	ErrOnImageIsInCollection bool
	ImageNotInCollection     bool
}

func (r FakeRepo) FindLabelByName(name string) (*lbl.Label, error) {
	if r.ErrOnFindLabel {
		return nil, r.Err
	}
	return lbl.NewLabel(name), nil
}

func (r FakeRepo) FindCollectionByName(name string) (*clc.Collection, error) {
	if r.ErrOnFindCollection {
		return nil, r.Err
	}
	return clc.NewCollection(name), nil
}

func (r FakeRepo) ImageIsInCollection(imageId im.ImageId, collectionId clc.CollectionId) (bool, error) {
	if r.ErrOnImageIsInCollection {
		return false, r.Err
	}
	if r.ImageNotInCollection {
		return false, nil
	}
	return true, nil
}
