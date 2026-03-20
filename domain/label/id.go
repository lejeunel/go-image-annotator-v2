package label

import (
	"github.com/google/uuid"
)

type LabelId uuid.UUID

func NewLabelID() LabelId {
	return LabelId(uuid.New())
}
