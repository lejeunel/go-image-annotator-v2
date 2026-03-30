package ingest

import (
	an "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type CollectionRepo interface {
	FindCollectionByName(string) (*clc.Collection, error)
}

type LabelRepo interface {
	FindLabelByName(string) (*lbl.Label, error)
}

type AnnotationRepo interface {
	AddLabelToImage(im.ImageId, clc.CollectionId, lbl.LabelId) error
	AddBoundingBoxToImage(im.ImageId, clc.CollectionId, an.BoundingBox) error
}

type ImageRepo interface {
	AddImage(im.ImageId, string) error
	AddImageToCollection(im.ImageId, clc.CollectionId) error
	FindImageIdByHash(string) (*im.ImageId, error)
	Delete(im.ImageId) error
}
