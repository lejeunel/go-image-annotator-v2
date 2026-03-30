package image

import (
	"github.com/google/uuid"
	uuidw "github.com/lejeunel/go-image-annotator-v2/uuid"
)

type ImageId struct {
	uuidw.UUIDWrapper[ImageId]
}

func NewImageId() ImageId {
	return ImageId{uuidw.UUIDWrapper[ImageId]{UUID: uuid.New()}}
}
