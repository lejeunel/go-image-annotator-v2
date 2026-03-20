package annotation

import (
	"github.com/google/uuid"
)

type AnnotationID uuid.UUID

func NewAnnotationID() AnnotationID {
	return AnnotationID(uuid.New())
}
