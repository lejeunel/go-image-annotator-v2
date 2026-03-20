package image

import (
	"github.com/google/uuid"
)

type ImageId uuid.UUID

func NewImageID() ImageId {
	return ImageId(uuid.New())
}
