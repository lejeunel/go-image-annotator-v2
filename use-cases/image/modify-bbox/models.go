package modify_bbox

import (
	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type Updatables struct {
	LabelId      lbl.LabelId
	AnnotationId a.AnnotationId
	Xc           float32
	Yc           float32
	Width        float32
	Height       float32
}

type Response struct{}

type Request struct {
	AnnotationId a.AnnotationId
	Label        string
	Xc           float32
	Yc           float32
	Width        float32
	Height       float32
}
