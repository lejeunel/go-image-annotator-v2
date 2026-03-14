package image

import (
	"github.com/google/uuid"
)

type ImageID uuid.UUID

func NewImageID() ImageID {
	return ImageID(uuid.New())
}
