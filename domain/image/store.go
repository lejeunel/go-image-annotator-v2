package image

import (
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type ImageStore interface {
	Find(BaseImage) (*Image, error)
}

type MyImageStore struct {
	repo         Repo
	artefactRepo ArtefactRepo
}

func (s *MyImageStore) Find(base BaseImage) (*Image, error) {
	collection, err := s.repo.FindCollectionByName(base.Collection)
	if err != nil {
		return nil, fmt.Errorf("fetching collection by name (%v): %w", base.Collection, err)
	}

	ok, err := s.repo.ImageExistsInCollection(base.ImageId, base.Collection)
	if err != nil {
		return nil, fmt.Errorf("checking whether image %v exists in collection %v: %w",
			base.ImageId.String(), base.Collection, err)
	}
	if !ok {
		return nil, fmt.Errorf("checking whether image %v exists in collection %v: %w",
			base.ImageId.String(), base.Collection, e.ErrNotFound)

	}

	labels, err := s.repo.FindLabels(base.ImageId, collection.Id)
	if err != nil {
		return nil, fmt.Errorf("fetching labels: %w", err)
	}

	boxes, err := s.repo.FindBoundingBoxes(base.ImageId, collection.Id)
	if err != nil {
		return nil, fmt.Errorf("fetching bounding boxes: %w", err)
	}

	return &Image{Id: base.ImageId, Collection: *collection, Labels: labels,
		BoundingBoxes: boxes,
		Reader:        *NewImageReader(base.ImageId, s.artefactRepo)}, nil

}

func NewImageStore(repo Repo, artefactRepo ArtefactRepo) *MyImageStore {
	return &MyImageStore{repo: repo, artefactRepo: artefactRepo}
}
