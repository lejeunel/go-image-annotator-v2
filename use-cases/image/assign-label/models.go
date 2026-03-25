package assign_label

import (
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type Response struct {
	ImageId    im.ImageId
	Collection string
	Label      string
}

type Request struct {
	ImageId    im.ImageId
	Collection string
	Label      string
}
