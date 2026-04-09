package image_store

import (
	"fmt"

	ast "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	rd "github.com/lejeunel/go-image-annotator-v2/application/image-reader"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type ImageStore struct {
	collectionRepo CollectionRepo
	annotationRepo AnnotationRepo
	imageRepo      ImageRepo
	artefactRepo   ast.ArtefactRepo
}

func (s ImageStore) Find(base im.BaseImage) (*im.Image, error) {
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

	labels, err := s.annotationRepo.FindImageLabels(base.ImageId, collection.Id)
	if err != nil {
		return nil, fmt.Errorf("fetching labels: %w", err)
	}

	boxes, err := s.annotationRepo.FindBoundingBoxes(base.ImageId, collection.Id)
	if err != nil {
		return nil, fmt.Errorf("fetching bounding boxes: %w", err)
	}

	return &im.Image{Id: base.ImageId, Collection: *collection, Labels: labels,
		BoundingBoxes: boxes,
		Reader:        rd.NewImageReader(base.ImageId, s.artefactRepo)}, nil

}

func NewImageStore(imageRepo ImageRepo, collectionRepo CollectionRepo,
	annotationRepo AnnotationRepo, artefactRepo ast.ArtefactRepo) *ImageStore {
	return &ImageStore{imageRepo: imageRepo,
		collectionRepo: collectionRepo,
		annotationRepo: annotationRepo,
		artefactRepo:   artefactRepo}
}
