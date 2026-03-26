package collection

import (
	"github.com/google/uuid"
)

type CollectionId uuid.UUID

func NewCollectionId() CollectionId {
	return CollectionId(uuid.New())
}

func (id CollectionId) String() string {
	return uuid.UUID(id).String()
}

func (id CollectionId) IsNil() bool {
	return uuid.UUID(id) == uuid.Nil
}
