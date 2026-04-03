package read_meta

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type OutputPort interface {
	Success(im.ImageResponse)
	ErrNotFound(error)
	ErrInternal(error)
}
