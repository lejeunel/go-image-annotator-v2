package image

import (
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type ImageStore interface {
	Find(BaseImage) (*Image, error)
}

type MyImageStore struct {
	collectionRepo CollectionRepo
	annotationRepo AnnotationRepo
	imageRepo      ImageRepo
	artefactRepo   ArtefactRepo
}

func (s *MyImageStore) Find(base BaseImage) (*Image, error) {
	collection, err := s.collectionRepo.FindCollectionByName(base.Collection)
	if err != nil {
		return nil, fmt.Errorf("fetching collection by name (%v): %w", base.Collection, err)
	}

	ok, err := s.imageRepo.ImageExistsInCollection(base.ImageId, collection.Id)
	if err != nil {
		return nil, fmt.Errorf("checking whether image %v exists in collection %v: %w",
			base.ImageId.String(), base.Collection, err)
	}
	if !ok {
		return nil, fmt.Errorf("checking whether image %v exists in collection %v: %w",
			base.ImageId.String(), base.Collection, e.ErrNotFound)

	}

	labels, err := s.annotationRepo.FindLabels(base.ImageId, collection.Id)
	if err != nil {
		return nil, fmt.Errorf("fetching labels: %w", err)
	}

	boxes, err := s.annotationRepo.FindBoundingBoxes(base.ImageId, collection.Id)
	if err != nil {
		return nil, fmt.Errorf("fetching bounding boxes: %w", err)
	}

	return &Image{Id: base.ImageId, Collection: *collection, Labels: labels,
		BoundingBoxes: boxes,
		Reader:        *NewImageReader(base.ImageId, s.artefactRepo)}, nil

}

func NewImageStore(imageRepo ImageRepo, collectionRepo CollectionRepo,
	annotationRepo AnnotationRepo, artefactRepo ArtefactRepo) *MyImageStore {
	return &MyImageStore{imageRepo: imageRepo,
		collectionRepo: collectionRepo,
		annotationRepo: annotationRepo,
		artefactRepo:   artefactRepo}
}
