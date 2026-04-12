package remove

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
)

type Repo interface {
	RemoveAnnotation(a.AnnotationId) error
}
