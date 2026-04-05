package delete

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type Repo interface {
	RemoveImageFromCollection(im.ImageId, clc.CollectionId) error
	RemoveAnnotation(im.ImageId, clc.CollectionId, a.AnnotationId) error
}
