package artefact

import (
	"github.com/google/uuid"
)

type ArtefactId uuid.UUID

func NewArtefactId() ArtefactId {
	return ArtefactId(uuid.New())
}
