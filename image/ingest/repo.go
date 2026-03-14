package ingest

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/artefact"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Repo interface {
	FindCollectionByName(string) (*clc.Collection, error)
	LabelExists(string) (bool, error)
	IngestImage(im.ImageID, clc.CollectionID, a.ArtefactID) error
}

type FakeRepo struct {
	GotImage            bool
	Err                 error
	ErrOnFindCollection bool
	ErrOnLabelExists    bool
	ErrOnIngest         bool
	CollectionExists_   bool
	LabelExists_        bool
}

func (r *FakeRepo) FindCollectionByName(name string) (*clc.Collection, error) {
	if r.ErrOnFindCollection {
		return nil, r.Err
	}
	if !r.CollectionExists_ {
		return nil, e.ErrNotFound
	}
	return clc.NewCollection("a-collection"), nil
}

func (r *FakeRepo) LabelExists(name string) (bool, error) {
	if r.ErrOnLabelExists {
		return false, r.Err
	}
	return r.LabelExists_, nil
}

func (r *FakeRepo) IngestImage(im.ImageID, clc.CollectionID, a.ArtefactID) error {
	return nil
}
