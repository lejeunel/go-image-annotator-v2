package add_bbox

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
)

type Repo interface {
	AddBoundingBox(im.ImageId, clc.CollectionId, a.BoundingBox) error
	FindLabelByName(string) (*lbl.Label, error)
}
