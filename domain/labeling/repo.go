package labeling

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type Repo interface {
	FindLabelIdByName(string) (*lbl.LabelId, error)
	FindCollectionIdByName(string) (*clc.CollectionId, error)
	ImageIsInCollection(im.ImageId, clc.CollectionId) (bool, error)
}

type FakeRepo struct {
	Err                      error
	ErrOnFindLabel           bool
	ErrOnFindCollection      bool
	ErrOnImageIsInCollection bool
	ImageNotInCollection     bool
}

func (r FakeRepo) FindLabelIdByName(name string) (*lbl.LabelId, error) {
	if r.ErrOnFindLabel {
		return nil, r.Err
	}
	return &lbl.LabelId{}, nil
}

func (r FakeRepo) FindCollectionIdByName(name string) (*clc.CollectionId, error) {
	if r.ErrOnFindCollection {
		return nil, r.Err
	}
	return &clc.CollectionId{}, nil
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
