package collection

import (
	"github.com/google/uuid"
)

type CollectionId uuid.UUID

func NewCollectionID() CollectionId {
	return CollectionId(uuid.New())
}
