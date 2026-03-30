package image_store

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type CollectionRepo interface {
	FindCollectionByName(string) (*clc.Collection, error)
}

type AnnotationRepo interface {
	FindImageLabels(im.ImageId, clc.CollectionId) ([]*a.ImageLabel, error)
	FindBoundingBoxes(im.ImageId, clc.CollectionId) ([]*a.BoundingBox, error)
}

type ImageRepo interface {
	ImageExistsInCollection(im.ImageId, clc.CollectionId) (bool, error)
}
