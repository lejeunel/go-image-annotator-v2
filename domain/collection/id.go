package collection

import (
	"github.com/google/uuid"
)

type CollectionID uuid.UUID

func NewCollectionID() CollectionID {
	return CollectionID(uuid.New())
}
