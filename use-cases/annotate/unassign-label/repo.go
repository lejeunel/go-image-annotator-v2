package unassign_label

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
)

type Repo interface {
	RemoveImageLabel(im.ImageId, clc.CollectionId, lbl.LabelId) error
}
