package read_raw

import (
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type Request struct {
	ImageId im.ImageId
}

type Response struct {
	Data []byte
}
