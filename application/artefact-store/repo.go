package artefact_store

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type ArtefactRepo interface {
	Store(im.ImageId, []byte) error
	Delete(im.ImageId) error
	Get(im.ImageId) ([]byte, error)
}

type ArtefactReadRepo interface {
	Get(im.ImageId) ([]byte, error)
}
