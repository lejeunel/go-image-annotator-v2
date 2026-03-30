package annotation

import (
	"github.com/google/uuid"
)

type AnnotationId uuid.UUID

func NewAnnotationId() AnnotationId {
	return AnnotationId(uuid.New())
}

func (id AnnotationId) String() string {
	return uuid.UUID(id).String()
}
