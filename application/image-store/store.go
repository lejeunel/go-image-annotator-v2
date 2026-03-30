package image_store

import (
	"fmt"

	ast "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type ImageStore interface {
	Find(im.BaseImage) (*im.Image, error)
}

type MyImageStore struct {
	collectionRepo CollectionRepo
	annotationRepo AnnotationRepo
	imageRepo      ImageRepo
	artefactRepo   ast.ArtefactRepo
}

func (s *MyImageStore) Find(base im.BaseImage) (*im.Image, error) {
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

	return &im.Image{Id: base.ImageId, Collection: *collection, Labels: labels,
		BoundingBoxes: boxes,
		Reader:        *ast.NewImageReader(base.ImageId, s.artefactRepo)}, nil

}

func NewImageStore(imageRepo ImageRepo, collectionRepo CollectionRepo,
	annotationRepo AnnotationRepo, artefactRepo ast.ArtefactRepo) *MyImageStore {
	return &MyImageStore{imageRepo: imageRepo,
		collectionRepo: collectionRepo,
		annotationRepo: annotationRepo,
		artefactRepo:   artefactRepo}
}
