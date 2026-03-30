package remove_bbox

import (
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
)

type Response struct{}

type Request struct {
	Id a.AnnotationId
}
