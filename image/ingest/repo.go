package ingest

import (
	an "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	a "github.com/lejeunel/go-image-annotator-v2/domain/artefact"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type Repo interface {
	FindCollectionByName(string) (*clc.Collection, error)
	FindLabelByName(string) (*lbl.Label, error)
	IngestImage(im.ImageId, clc.CollectionId, a.ArtefactId) error
	AddLabelToImage(im.ImageId, clc.CollectionId, lbl.LabelId) error
	AddBoundingBoxToImage(im.ImageId, clc.CollectionId, an.BoundingBox) error
	FindImageByHash(string) (*im.Image, error)
	DeleteImage(im.ImageId) error
}
