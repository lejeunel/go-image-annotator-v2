package annotation

import (
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type ImageLabel struct {
	Id    AnnotationId
	Label lbl.Label
}

type BoundingBox struct {
	Id     AnnotationId
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
	return &BoundingBox{Id: NewAnnotationId(), Xc: xc, Yc: yc, Width: width, Height: height, Label: label}
}

func NewImageLabel(label lbl.Label) *ImageLabel {
	return &ImageLabel{Id: NewAnnotationId(), Label: label}
}
