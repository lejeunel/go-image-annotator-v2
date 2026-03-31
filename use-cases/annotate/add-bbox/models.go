package add_bbox

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type Response struct{}

type Request struct {
	ImageId    im.ImageId
	Collection string
	Label      string
	Xc         float32
	Yc         float32
	Width      float32
	Height     float32
}
