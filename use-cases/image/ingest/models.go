package ingest

import (
	an "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type Request struct {
	Collection    string
	Labels        []string
	BoundingBoxes []BoundingBoxRequest
	Reader        im.ImageReader
}

type Response struct {
	ImageId       im.ImageId
	Collection    string
	Labels        []string
	BoundingBoxes []an.BoundingBoxResponse
}

func NewIngestionResponse(image *im.Image) Response {
	return Response{ImageId: image.Id, Collection: image.Collection.Name, Labels: image.LabelNames(),
		BoundingBoxes: image.BoundingBoxSummary()}
}

type BoundingBoxRequest struct {
	Label  string
	Xc     float32
	Yc     float32
	Width  float32
	Height float32
}
