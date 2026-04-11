package read_raw

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"io"
)

type Request struct {
	ImageId im.ImageId
}

type Response struct {
	Reader io.Reader
}
