package import_shallow

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type Repo interface {
	ImageExists(im.ImageId) (bool, error)
	FindCollection(string) (*clc.Collection, error)
	ImageExistsInCollection(im.ImageId, clc.CollectionId) (bool, error)
	ImportImage(im.ImageId, clc.CollectionId) error
}
