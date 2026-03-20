package annotation

import (
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type BoundingBox struct {
	ID     AnnotationID
	Label  lbl.Label
	Xc     float32
	Yc     float32
	Width  float32
	Height float32
}

type BoundingBoxResponse struct {
	Label  string
	Xc     float32
	Yc     float32
	Width  float32
	Height float32
}

func NewBoundingBox(xc float32, yc float32, width float32, height float32, label lbl.Label) *BoundingBox {
	return &BoundingBox{ID: NewAnnotationID(), Xc: xc, Yc: yc, Width: width, Height: height, Label: label}
}
