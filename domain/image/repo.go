package image

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
)

type CollectionRepo interface {
	FindCollectionByName(string) (*clc.Collection, error)
}

type AnnotationRepo interface {
	FindLabels(ImageId, clc.CollectionId) ([]*a.ImageLabel, error)
	FindBoundingBoxes(ImageId, clc.CollectionId) ([]*a.BoundingBox, error)
}

type ImageRepo interface {
	ImageExistsInCollection(ImageId, clc.CollectionId) (bool, error)
}
