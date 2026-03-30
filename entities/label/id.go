package label

import (
	"github.com/google/uuid"
	uuidw "github.com/lejeunel/go-image-annotator-v2/uuid"
)

type LabelId struct{ uuidw.UUIDWrapper[LabelId] }

func NewLabelId() LabelId {
	return LabelId{uuidw.UUIDWrapper[LabelId]{UUID: uuid.New()}}
}
