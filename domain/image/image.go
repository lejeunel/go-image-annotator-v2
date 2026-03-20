package image

import (
	"fmt"
	an "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	a "github.com/lejeunel/go-image-annotator-v2/domain/artefact"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type RawImage struct {
	ArtefactId a.ArtefactId
	Data       []byte
	Hash       string
}

type Image struct {
	Id            ImageId
	ArtefactId    a.ArtefactId
	Collection    clc.Collection
	Hash          string
	Labels        []*lbl.Label
	BoundingBoxes []*an.BoundingBox
}

func (i *Image) AddLabel(l *lbl.Label) error {
	for _, label := range i.Labels {
		if l.Name == label.Name {
			return fmt.Errorf("adding label %v to image: found duplicate: %w", label.Name, e.ErrValidation)
		}
	}
	i.Labels = append(i.Labels, l)
	return nil
}

func (i *Image) AddBoundingBox(xc float32, yc float32, width float32, height float32, label lbl.Label) error {
	errCtx := "adding bounding box to image"
	if width <= 0 {
		return fmt.Errorf("%v: checking whether width (%v) <= 0: %w", errCtx, width, e.ErrValidation)
	}
	i.BoundingBoxes = append(i.BoundingBoxes, an.NewBoundingBox(xc, yc, width, height, label))
	return nil
}

func (i *Image) LabelNames() []string {
	labelNames := []string{}
	for _, label := range i.Labels {
		labelNames = append(labelNames, label.Name)
	}
	return labelNames
}

func (i *Image) BoundingBoxSummary() []an.BoundingBoxResponse {
	summary := []an.BoundingBoxResponse{}
	for _, bbox := range i.BoundingBoxes {
		summary = append(summary,
			an.BoundingBoxResponse{Label: bbox.Label.Name, Xc: bbox.Xc, Yc: bbox.Yc, Width: bbox.Width, Height: bbox.Height})
	}
	return summary
}

func NewImage(hash string, collection clc.Collection, artefactId a.ArtefactId) *Image {
	return &Image{Id: NewImageID(), Collection: collection, Hash: hash, ArtefactId: artefactId}
}
