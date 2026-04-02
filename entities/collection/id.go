package collection

import (
	"github.com/google/uuid"
	uuidw "github.com/lejeunel/go-image-annotator-v2/shared/uuid"
)

type CollectionId struct {
	uuidw.UUIDWrapper[CollectionId]
}

func NewCollectionId() CollectionId {
	return CollectionId{uuidw.UUIDWrapper[CollectionId]{UUID: uuid.New()}}
}
