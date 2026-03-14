package artefact

import (
	"github.com/google/uuid"
)

type ArtefactID uuid.UUID

func NewArtefactID() ArtefactID {
	return ArtefactID(uuid.New())
}
