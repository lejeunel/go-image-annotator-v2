package image

import (
	"github.com/google/uuid"
)

type ImageId uuid.UUID

func (id ImageId) String() string {
	return uuid.UUID(id).String()
}
func NewImageId() ImageId {
	return ImageId(uuid.New())
}
