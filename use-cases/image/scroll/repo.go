package scroll

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type Repo interface {
	GetAdjacent(im.ImageId, string, ScrollingDirection) (*im.BaseImage, error)
}
