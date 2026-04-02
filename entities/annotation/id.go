package annotation

import (
	"github.com/google/uuid"
	uuidw "github.com/lejeunel/go-image-annotator-v2/shared/uuid"
)

type AnnotationId struct {
	uuidw.UUIDWrapper[AnnotationId]
}

func NewAnnotationId() AnnotationId {
	return AnnotationId{uuidw.UUIDWrapper[AnnotationId]{UUID: uuid.New()}}
}
