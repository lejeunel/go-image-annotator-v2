package read_meta

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type Request struct {
	ImageId    im.ImageId
	Collection string
}

type Response struct {
	Id         im.ImageId
	Collection string
	Labels     []*a.ImageLabel
}
