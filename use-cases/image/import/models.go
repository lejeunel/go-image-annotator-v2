package import_image

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type Request struct {
	ImageId    im.ImageId
	Collection string
}

type Response struct{}
